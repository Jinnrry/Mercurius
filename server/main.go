package main

import (
	"Mercurius/common"
	"Mercurius/server/worker"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
)

var configPath string
var showVersion bool
var config common.Config

// 日志初始化
func init() {
	// 初始化配置文件
	flag.StringVar(&configPath, "c", "./config.json", "config文件位置")
	flag.BoolVar(&showVersion, "v", false, "显示程序版本")
	flag.Parse()

	if showVersion {
		common.PrintVersion("Mercurius Server")
		os.Exit(0)
	}

	var err error
	config, err = common.GetConfig(configPath)

	if err != nil {
		log.Fatalf("配置文件加载错误:%v", err)
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	serverStart(config)
}

func serverStart(config common.Config) {
	// 启动worker
	for idx, item := range config.Client.Services {
		go worker.CreateWorker(idx, item.RemotPort)
	}

	// 启动master
	master := worker.GetMasterInstance()
	err := master.Run(config.Server.Port)
	if err != nil {
		log.Fatalf("程序启动失败:%v", err)
	}

}
