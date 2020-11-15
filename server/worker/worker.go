package worker

import (
	"Mercurius/common"
	"fmt"
	"log"
	"net"
	"sync"
)

// socker池
var RequestConnPool sync.Map

func CreateWorker(serviceId int, port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error listen:%v", err)
		return
	}
	defer listen.Close()
	for {
		if requestConn, err := listen.Accept(); err == nil {
			log.Println("Request 接入")

			go requestHandle(serviceId, requestConn)
		} else {
			log.Fatalf("accept error:%v", err)
		}
	}

}

func CloseRequestSocket(poolKey string) {
	conn, ok := RequestConnPool.Load(poolKey)
	if ok {
		(conn.(net.Conn)).Close()
		RequestConnPool.Delete(poolKey)
	}
}

// 接收request来的数据
func requestHandle(serviceId int, requestConn net.Conn) {
	requestId := common.GetId()

	// 通知client打开连接
	SendOpenCommand(requestId, uint16(serviceId))

	poolKey := fmt.Sprintf("%d_%d", serviceId, requestId)
	RequestConnPool.Store(poolKey, requestConn)

	buf := make([]byte, common.TransmissionDataLength)
	//循环读取RequestClient数据流
	for {
		//网络数据流读入 buffer
		cnt, err := requestConn.Read(buf)
		//数据读尽、读取错误 关闭 socket 连接
		if cnt == 0 || err != nil {
			log.Println("request socket断开")

			requestConn.Close()
			RequestConnPool.Delete(poolKey)

			// 通知client断开socket
			SendCloseCommand(requestId, uint16(serviceId))

			break
		}

		// 将读取到的buffer写到Mercurius Client
		data := common.TransmissionStruct{
			RequestId:  requestId,
			Data:       buf,
			DataLength: uint16(cnt),
			ServiceId:  uint16(serviceId),
		}
		_, err = GetMasterInstance().SendData2Client(data)
		// 写入client错误，关闭调request socket
		if err != nil {
			log.Println("request socket断开")
			requestConn.Close()
			RequestConnPool.Delete(poolKey)
			SendCloseCommand(requestId, uint16(serviceId))
		}

	}
}

func Senddata2Request(data common.TransmissionStruct) {

	poolKey := fmt.Sprintf("%d_%d", data.ServiceId, data.RequestId)
	conn, ok := RequestConnPool.Load(poolKey)
	if ok {
		log.Println("发送数据给Request")
		(conn.(net.Conn)).Write(data.GetData())
	}
}

// 通知client断开本地socket
func SendCloseCommand(requestId uint64, serviceId uint16) {
	data := []byte(fmt.Sprintf("%d,%d", requestId, serviceId))
	tdata := common.TransmissionStruct{
		RequestId:  common.CloseSocket,
		ServiceId:  common.CloseSocket,
		DataLength: uint16(len(data)),
	}

	paddingLen := common.TransmissionDataLength - len(data)

	padding := make([]byte, paddingLen)

	tdata.Data = append(data, padding...)

	GetMasterInstance().SendData2Client(tdata)
}

// 通知client打开本地socket
func SendOpenCommand(requestId uint64, serviceId uint16) {
	data := []byte(fmt.Sprintf("%d,%d", requestId, serviceId))
	tdata := common.TransmissionStruct{
		RequestId:  common.OpenSocket,
		ServiceId:  serviceId,
		DataLength: uint16(len(data)),
	}

	paddingLen := common.TransmissionDataLength - len(data)

	padding := make([]byte, paddingLen)

	tdata.Data = append(data, padding...)

	GetMasterInstance().SendData2Client(tdata)
}
