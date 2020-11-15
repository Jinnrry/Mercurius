package worker

import (
	"Mercurius/common"
	"net"
	"sync"
)

type Master interface {
	Run(int) error                                          // 启动Master
	clientDataHandle(net.Conn)                              // 接收处理client的数据
	SendData2Client(common.TransmissionStruct) (int, error) // 发送数据给client
}

var instance Master
var mu sync.Mutex

func GetMasterInstance() Master {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &TcpMaster{
			verify: false,
		}
	}
	return instance

}
