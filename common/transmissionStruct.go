package common

import (
	"Mercurius/common/encryption"
	"encoding/binary"
	log "github.com/sirupsen/logrus"
)

type TransmissionStruct struct {
	RequestId  uint64
	ServiceId  uint16
	DataLength uint16
	Data       []byte
}

const (
	CloseSocket = 1 // 断开本地socket通知
	OpenSocket  = 3 // 打开本地socket通知
	VerifyKey   = 2 // 验证客户端秘钥

	TransmissionPackageLength = 4112 // 8字节的RequestID， 2字节的serviceid  2字节dataLength  4096字节的数据包长度
	TransmissionDataLength    = 4096 // 传输的数据包长度
)

func (t *TransmissionStruct) GetData() []byte {
	return t.Data[:t.DataLength]
}

func (t *TransmissionStruct) Convert2Byte() []byte {
	requestId := make([]byte, 8)
	binary.LittleEndian.PutUint64(requestId, uint64(t.RequestId))

	serviceId := make([]byte, 2)
	binary.LittleEndian.PutUint16(serviceId, uint16(t.ServiceId))

	dataLength := make([]byte, 2)
	binary.LittleEndian.PutUint16(dataLength, uint16(t.DataLength))

	ret := append(requestId, serviceId...)
	ret = append(ret, dataLength...)
	ret = append(ret, t.Data...)

	// 校验数据包长度
	if len(t.Data) < TransmissionDataLength {
		padding := make([]byte, TransmissionDataLength-len(t.Data))
		ret = append(ret, padding...)
	}

	// 数据包加密
	config, _ = GetConfig("")
	encrData, _ := encryption.Encrypt(ret, []byte(config.Common.Token))

	return encrData
}

func FactoryTransmission(encrData []byte) TransmissionStruct {
	// 数据包解密
	config, _ = GetConfig("")
	data, err := encryption.Decrypt(encrData, []byte(config.Common.Token))
	if err != nil {
		log.Errorf("数据包解密失败！ %v", err)
	}
	requestId := binary.LittleEndian.Uint64(data[0:8])
	serviceId := binary.LittleEndian.Uint16(data[8:10])
	dataLength := binary.LittleEndian.Uint16(data[10:12])

	tdata := data[12:]
	ret := TransmissionStruct{
		RequestId:  requestId,
		ServiceId:  serviceId,
		DataLength: dataLength,
		Data:       tdata,
	}

	return ret
}
