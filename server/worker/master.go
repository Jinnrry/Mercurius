package worker

import (
	"Mercurius/common"
	"sync"
)

type Master interface {
	Run(int) error                                          // 启动Master
	SendData2Client(common.TransmissionStruct) (int, error) // 发送数据给client
}

var instance Master
var mu sync.Mutex

func GetMasterInstance() Master {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		config := common.GetConfig()
		if config.Common.Protocol == "tcp" {
			instance = &TcpMaster{
				verify: false,
			}
		} else {
			instance = &WebSocketMaster{}
		}

	}
	return instance

}
