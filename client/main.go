package main

import (
	"Mercurius/client/worker"
	"Mercurius/common"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
)

// 初始化配置文件
var configPath string
var showVersion bool
var config common.Config

// 资源初始化
func init() {
	flag.StringVar(&configPath, "c", "./config.json", "config文件位置")
	flag.BoolVar(&showVersion, "v", false, "显示程序版本")
	flag.Parse()

	if showVersion {
		common.PrintVersion("Mercurius Client")
		os.Exit(0)
	}

	// 配置初始化
	var err error
	config, err = common.InitConfig(configPath)
	if err != nil {
		log.Fatalf("配置文件加载错误:%v", err)
	}

	//日志配置初始化
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debugf("配置详情%+v", config)
	clientStart(config)
}

func clientStart(config common.Config) {
	clientWorker := worker.GetClientWorkerInstance()

	err := clientWorker.Run(config)
	if err != nil {
		log.Fatalf("程序启动失败:%v", err)
	}
}
