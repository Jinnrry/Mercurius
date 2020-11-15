package common

import (
	"fmt"
	"testing"
)

func TestTransmissionStruct_Conver2Byte(t1 *testing.T) {
	t := TransmissionStruct{
		RequestId: 9999,
		Data:      []byte{1, 0, 11, 2},
		ServiceId: 9,
	}
	fmt.Println(t.Convert2Byte())
	fmt.Println(FactoryTransmission(t.Convert2Byte()))
}
