package main

import (
	"Mercurius/client/worker"
	"Mercurius/common"
	"flag"
	"log"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "c", "./config.json", "config文件位置,默认./config.json")
	flag.Parse()
	config, err := common.GetConfig(configPath)
	if err != nil {
		log.Fatal("配置文件加载错误:%v", err)
	}
	clientWorker := worker.GetClientWorkerInstance()

	err = clientWorker.Run(config)
	if err != nil {
		log.Fatal("程序启动失败:%v", err)
	}
}
