package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"wukong/pkg/config"
)

func initConfig(cmd *cobra.Command) error {
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

	// 设置默认配置文件
	if len(configFile) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		configFile = dir + "/config.yaml"
		configFile = filepath.Clean(configFile)
		if _, err := os.Stat(configFile); err != nil {
			configFile = ""
		}
	}

	// 读取配置文件
	if len(configFile) > 0 {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// 绑定命令行参数到配置项
	// 配置项优先级：命令行参数 > 配置文件 > 默认命令行参数
	_ = viper.BindPFlags(cmd.Flags())
	_ = viper.BindPFlag("server_bind", cmd.Flags().Lookup("server-bind"))
	_ = viper.BindPFlag("config_secret_key", cmd.Flags().Lookup("config-secret-key"))
	_ = viper.BindPFlag("log_level", cmd.Flags().Lookup("log-level"))

	err := viper.Unmarshal(&config.Conf)
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

	if len(configFile) == 0 {
		logrus.Warn("no specified config file, maybe should use '-c config.yaml' flag")
	}

	if viper.GetBool("pprof-server") {
		go pprofServer(7777)
	}

	logrus.Debugf("config init completed: %+v", string(xutil.RemoveError(json.Marshal(config.Conf))))
	return nil
}

func pprofServer(port int) {
	logrus.Infof("pprof server listening on 0.0.0.0:%d", port)
	err := http.ListenAndServe("0.0.0.0:"+fmt.Sprintf("%d", port), nil)
	if err != nil {
		logrus.Warnf("pprof server error: %s", err)
		pprofServer(port + 1)
	}
}
