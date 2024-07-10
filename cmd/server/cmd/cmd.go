package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"wukong/pkg/config"
	"wukong/pkg/configpost"
)

var cmdReady bool

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Wukong Server",
	Run: func(cmd *cobra.Command, args []string) {
		cmdReady = true
	},
	Example: "server -c config.yaml",
}

var configFile string

const (
	logLevel   = 4
	serverBind = "0.0.0.0:8080"
)

func init() {
	// 初始化命令行参数
	defaultConfig := "./config/config.yaml"
	if os.Getenv("ENV_NAME") != "" {
		defaultConfig = "./config/config-" + os.Getenv("ENV_NAME") + ".yaml"
	}
	rootCmd.Flags().StringVarP(&configFile, "config", "c", defaultConfig, "config file")
	rootCmd.Flags().IntP("log-level", "l", logLevel, "log level")
	rootCmd.Flags().StringP("server-bind", "s", serverBind, "server bind address")
	rootCmd.Flags().StringP("config-secret-key", "", "", "config secret key")
}

func initConfig() error {
	// 设置日志格式
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})

	viper.SetConfigName("config")

	// 配置文件和命令行参数都不指定时的默认配置
	// viper.SetDefault("conn_timeout", 10)

	// 读取配置文件
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// 绑定命令行参数到配置项
	// 配置项优先级：命令行参数 > 配置文件 > 默认命令行参数
	_ = viper.BindPFlag("server_bind", rootCmd.Flags().Lookup("server-bind"))
	_ = viper.BindPFlag("config_secret_key", rootCmd.Flags().Lookup("config-secret-key"))
	_ = viper.BindPFlag("log_level", rootCmd.Flags().Lookup("log-level"))

	err = viper.Unmarshal(&config.Conf)
	if err != nil {
		return err
	}

	if len(config.Conf.ConfigSecretKey) == 0 {
		config.Conf.ConfigSecretKey = os.Getenv("CONFIG_SECRET_KEY")
	}

	logrus.SetLevel(logrus.Level(config.Conf.LogLevel))
	gin.SetMode(gin.ReleaseMode)
	if logrus.Level(config.Conf.LogLevel) >= logrus.DebugLevel {
		gin.SetMode(gin.DebugMode)
		logrus.SetReportCaller(true)
	}
	return nil
}

func Execute(ctx context.Context) {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("cmd error: %s", err.Error())
	}
	if cmdReady {
		err := initConfig()
		if err != nil {
			logrus.Fatalf("config error: %s", err.Error())
		}
		initConfigPost(ctx)
	}
}

func initConfigPost(ctx context.Context) {
	configpost.PostInit(ctx)
}
