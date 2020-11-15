package main

import (
	"Mercurius/common"
	"Mercurius/server/worker"
	"flag"
	"log"
)

func main() {
	// 初始化配置文件

	var configPath string
	flag.StringVar(&configPath, "c", "./config.json", "config文件位置,默认./config.json")
	flag.Parse()

	config, err := common.GetConfig(configPath)

	if err != nil {
		log.Fatal("配置文件加载错误:%v", err)
	}

	// 启动worker
	for idx, item := range config.Client.Services {
		go worker.CreateWorker(idx, item.RemotPort)
	}

	// 启动master
	master := worker.GetMasterInstance()
	err = master.Run(config.Server.Port)
	if err != nil {
		log.Fatal("程序启动失败:%v", err)
	}

}
