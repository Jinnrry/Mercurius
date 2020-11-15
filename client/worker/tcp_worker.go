package worker

import (
	"Mercurius/common"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type TcpWorker struct {
	serverConn net.Conn
	verify     bool
}

func (t *TcpWorker) Run(config common.Config) error {
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port))
		if err != nil {
			log.Printf("与服务端连接失败，10秒后重试,错误信息： %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		t.serverConn = conn
		// tcp连接成功后发送权限包

		t.SendData2Server(common.TransmissionStruct{
			RequestId:  common.VerifyKey,
			ServiceId:  0,
			Data:       []byte(config.Common.Token),
			DataLength: uint16(len([]byte(config.Common.Token))),
		})

		t.serverDataHandle(config, conn)
		log.Println("与服务端连接断开，10秒后重试")
		time.Sleep(10 * time.Second)
	}

	return nil
}

func (t *TcpWorker) SendData2Server(transmissionStruct common.TransmissionStruct) {
	log.Println("发送数据给Server")

	t.serverConn.Write(transmissionStruct.Convert2Byte())
}

// 读取Mercurius Server的数据
func (t *TcpWorker) serverDataHandle(config common.Config, conn net.Conn) {

	//缓存 conn 中的数据
	buf := make([]byte, common.TransmissionPackageLength)

	for {

		cnt, err := conn.Read(buf)

		if cnt == 0 || err != nil {
			log.Println("与服务端连接断开")
			conn.Close()
			break
		}

		pkg := common.FactoryTransmission(buf)
		log.Printf("Requestid %d", pkg.RequestId)

		if pkg.RequestId == common.CloseSocket {
			// 处理系统消息
			data := string(pkg.GetData())
			log.Println("关闭本地socket", data)

			info := strings.Split(data, ",")

			poolKey := fmt.Sprintf("%s_%s", info[1], info[0])
			pool_conn, ok := mercuriusClientConnPool.Load(poolKey)

			if ok {
				(pool_conn.(net.Conn)).Close()
				mercuriusClientConnPool.Delete(poolKey)
			}

			continue
		}

		if pkg.RequestId == common.OpenSocket {
			// 处理系统消息
			data := string(pkg.GetData())

			info := strings.Split(data, ",") // requestid , serviceid

			requestId, _ := strconv.ParseUint(info[0], 10, 64)
			serviceId, _ := strconv.ParseUint(info[1], 10, 16)
			poolKey := fmt.Sprintf("%s_%s", info[1], info[0])
			_, ok := mercuriusClientConnPool.Load(poolKey)

			// 已有就不处理了
			if ok {
				continue
			} else {
				createLocalSocket(config, requestId, uint16(serviceId), poolKey)
			}
			continue
		}

		log.Println("接收到Server数据")
		//log.Println(string(pkg.Data))

		poolKey := fmt.Sprintf("%d_%d", pkg.ServiceId, pkg.RequestId)

		pool_conn, ok := mercuriusClientConnPool.Load(poolKey)
		if !ok {
			// 没有本地连接的时候先创建连接
			log.Printf("创建位置2")
			pool_conn = createLocalSocket(config, pkg.RequestId, pkg.ServiceId, poolKey)
		}

		_, err = (pool_conn.(net.Conn)).Write(pkg.GetData())
		if err != nil {
			log.Printf("写入本地socket错误 %v", err)
			(pool_conn.(net.Conn)).Close()
			mercuriusClientConnPool.Delete(poolKey)
		}

	}
}
