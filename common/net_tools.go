package common

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

func GetDataFromConn(conn net.Conn, dataLength int) ([]byte, error) {

	totalDataLen := 0
	totalData := make([]byte, dataLength)

	for {
		//读取 conn 中的数据
		buf := make([]byte, dataLength-totalDataLen)
		cnt, err := conn.Read(buf)
		log.Debugf("读取数据长度 %d", cnt)

		if cnt == 0 || err != nil {
			log.Info("与服务端连接断开")
			return nil, errors.New(fmt.Sprintf("与服务端通讯错误:%v", err))
		}
		totalData = append(totalData[:totalDataLen], buf...)
		totalDataLen += cnt
		if totalDataLen == dataLength {
			break
		}
	}

	return totalData, nil
}
