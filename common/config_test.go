package common

import (
	"testing"
)

func TestGetConfig(t *testing.T) {

	res, err := GetConfig("../config.json")

	if err != nil {
		t.Errorf("配置文件读取错误:%v", err)
	} else if res.Common.Token == "" {
		t.Errorf("配置文件格式错误")
	} else {
		t.Logf("配置文件测试通过，配置详情：\n%+v", res)
	}
}
