package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

type Config struct {
	Common struct {
		Token string `json:"token"`
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

func GetConfig(configPath string) (Config, error) {
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
		return config, err

	} else {
		return config, nil
	}

}
