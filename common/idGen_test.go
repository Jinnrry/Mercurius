package common

import (
	"fmt"
	"testing"
	"time"
)

func TestGetId(t *testing.T) {

	// 100个协程测试并发id生成
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(GetId())
		}()

	}

	time.Sleep(1 * time.Second)

}
