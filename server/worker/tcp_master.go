package worker

import (
	"Mercurius/common"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type TcpMaster struct {
	mercuriusClientConn net.Conn // 与client通讯的tcp socket
	verify              bool     // 权限验证是否通过
}

func (m *TcpMaster) Run(port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error listen:%v", err)
		return err
	}

	log.Println("Mercurius 启动成功！")

	defer listen.Close()
	for {
		if conn, err := listen.Accept(); err != nil {
			log.Fatalf("accept error:%v", err)

		} else {
			log.Println("客户端接入")
			// 设置一个3秒的计时器，3秒后没通过权限验证则关闭连接
			go func() {
				log.Println("权限验证计时器")
				time.Sleep(3 * time.Second)
				if !m.verify {
					log.Println("权限验证超时，关闭连接")
					conn.Close()
				}
			}()
			m.mercuriusClientConn = conn
			m.clientDataHandle(conn)
		}
	}

}

// 接收client来的数据
func (m *TcpMaster) clientDataHandle(conn net.Conn) {
	buf := make([]byte, common.TransmissionPackageLength)

	//循环读取RequestClient数据流
	for {
		//网络数据流读入 buffer
		cnt, err := conn.Read(buf)
		//数据读尽、读取错误 关闭 socket 连接
		if cnt == 0 || err != nil {
			log.Println("客户端连接关闭")
			m.verify = false
			conn.Close()
			break
		}

		dataPackage := common.FactoryTransmission(buf)
		log.Println("服务端接收到客户端数据包")

		if !m.verify && dataPackage.RequestId != common.VerifyKey {
			// 验证未通过切不是权限验证的数据包
			conn.Close()
			m.verify = false
			log.Println("权限校验失败！关闭连接")
			break
		}

		if dataPackage.RequestId == common.VerifyKey {
			config, _ := common.GetConfig("")
			if string(dataPackage.GetData()) == config.Common.Token {
				log.Println("权限校验通过")
				m.verify = true
			} else {
				conn.Close()
				m.verify = false
				log.Println("权限校验失败！关闭连接")
				break
			}
		}

		if dataPackage.RequestId == common.CloseSocket {
			// 处理系统消息
			data := string(dataPackage.GetData())
			log.Println("关闭Request Socket", data)
			info := strings.Split(data, ",") // request_id  serviceid
			poolKey := fmt.Sprintf("%d_%d", info[1], info[0])
			CloseRequestSocket(poolKey)
			continue
		}

		// 将数据包发送给request
		Senddata2Request(dataPackage)

	}

}

// 将数据发送到Mercurius客户端
func (m *TcpMaster) SendData2Client(data common.TransmissionStruct) (int, error) {
	if m.mercuriusClientConn != nil && m.verify {
		log.Println("发送数据到Client")
		return m.mercuriusClientConn.Write(data.Convert2Byte())
	}
	return 0, errors.New("客户端未连接")
}
