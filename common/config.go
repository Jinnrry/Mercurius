package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Config struct {
	Common struct {
		Token    string `json:"token"`
		Protocol string `json:"protocol"`
	} `json:"common"`
	Server struct {
		Port int    `json:"port"`
		Ip   string `json:"ip"`
	} `json:"server"`
	Client struct {
		Services []struct {
			LocalIp   string `json:"local_ip"`
			LocalPort int    `json:"local_port"`
			RemotPort int    `json:"remot_port"`
			Type      string `json:"type"`
		} `json:"services"`
	} `json:"client"`
}

var config Config

var AvailableProtocol = map[string]string{
	"tcp":       "enable",
	"websocket": "enable",
}

func InitConfig(configPath string) (Config, error) {
	if reflect.DeepEqual(config, Config{}) {
		file, err := os.Open(configPath)
		if err != nil {
			return config, err
		}

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			return config, err
		}
		err = json.Unmarshal(fileContent, &config)

		if _, ok := AvailableProtocol[config.Common.Protocol]; !ok {
			panic(fmt.Sprintf("%s协议不支持！,目前仅支持websocket、tcp", config.Common.Protocol))
		}

		return config, err

	} else {
		return config, nil
	}

}

func GetConfig() Config {
	if reflect.DeepEqual(config, Config{}) {
		panic("未初始化配置文件")
	}

	return config
}
