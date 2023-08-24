package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/dao/secret"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/server/controller"
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
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})

	viper.SetConfigName("config")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
	if cmdReady {
		err := initConfig()
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Debugf("config %+v", config.Conf)

		initConfigPost()
	}
}

func initConfigPost() {
	err := secret.DecryptCFBToStruct(&config.Conf.Secret, config.Conf.ConfigSecretKey)
	if err != nil {
		logrus.Fatal(err)
	}

	db.InitRedis()
	db.InitMysql()
	controller.RunServer(config.Conf.ServerBind)
}
