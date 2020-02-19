package main

import (
	"os"

	glog "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitLogging initializes the logger based on the config
func InitLogging() {
	logFile := viper.GetString("log-file")
	logFormat := viper.GetString("log-format")
	logLevel := viper.GetString("log-level")

	switch logFile {
	case "stdout":
		glog.SetOutput(os.Stdout)
	case "stderr":
		glog.SetOutput(os.Stderr)
	default:
		file, err := os.Create(logFile)
		if err != nil {
			glog.Warnf("Couldn't open log-file '%s', falling back to stdout: %s", logFile, err)
			glog.SetOutput(os.Stdout)
		} else {
			glog.SetOutput(file)
		}

	}

	switch logFormat {
	case "text":
		glog.SetFormatter(&glog.TextFormatter{})
	case "json":
		glog.SetFormatter(&glog.JSONFormatter{})
	default:
		glog.Warnf("Unknown log-format '%s', falling back to 'text' format.", logFormat)
		glog.SetFormatter(&glog.TextFormatter{})
	}

	switch logLevel {
	case "debug":
		glog.SetLevel(glog.DebugLevel)
	case "info":
		glog.SetLevel(glog.InfoLevel)
	case "warning":
		glog.SetLevel(glog.WarnLevel)
	case "error":
		glog.SetLevel(glog.ErrorLevel)
	default:
		glog.Warnf("Unknown log-level '%s', falling back to 'warning' level.", logLevel)
		glog.SetLevel(glog.WarnLevel)
	}
}
