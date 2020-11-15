package worker

import (
	"Mercurius/common"
	"fmt"
	"log"
	"net"
	"sync"
)

var mu sync.Mutex
var instance ClientWorker

type ClientWorker interface {
	Run(config common.Config) error
	SendData2Server(data common.TransmissionStruct)
	serverDataHandle(config common.Config, conn net.Conn)
}

func GetClientWorkerInstance() ClientWorker {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &TcpWorker{
			verify: true,
		}
	}
	return instance
}

// 本地socket 连接池
var mercuriusClientConnPool sync.Map

// 读取本地socket数据
func connHandleLocal(requestId uint64, serviceId uint16, conn net.Conn) {

	//缓存 conn 中的数据
	buf := make([]byte, common.TransmissionDataLength)

	for {

		cnt, err := conn.Read(buf)

		if cnt == 0 || err != nil {
			conn.Close()
			poolKey := fmt.Sprintf("%d_%d", serviceId, requestId)
			mercuriusClientConnPool.Delete(poolKey)

			data := []byte(fmt.Sprintf("%d,%d\n", requestId, serviceId))

			log.Printf("通知服务端断开socket %d %d", requestId, serviceId)
			// 通知服务端断开socket
			GetClientWorkerInstance().SendData2Server(common.TransmissionStruct{
				RequestId:  common.CloseSocket,
				ServiceId:  serviceId,
				Data:       data,
				DataLength: uint16(len(data)),
			})

			break
		}

		pkg := common.TransmissionStruct{
			RequestId:  requestId,
			ServiceId:  serviceId,
			DataLength: uint16(cnt),
			Data:       buf,
		}

		log.Println("收到本地Socket数据,长度%d", cnt)

		// 发送给server
		GetClientWorkerInstance().SendData2Server(pkg)

	}
}

func createLocalSocket(config common.Config, requestId uint64, serverId uint16, poolKey string) net.Conn {
	// 没有连接的时候先创建连接
	local_conn, err := net.Dial("tcp",
		fmt.Sprintf("%s:%d", config.Client.Services[serverId].LocalIp, config.Client.Services[serverId].LocalPort))
	if err == nil {
		log.Println("创建本地连接成功")
		mercuriusClientConnPool.Store(poolKey, local_conn)
		go connHandleLocal(requestId, serverId, local_conn)
	} else {
		log.Fatalf("本地连接错误:%v", err)
	}

	return local_conn
}
