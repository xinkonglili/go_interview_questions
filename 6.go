/*
// 提供的公共函数
func SyncData() error
// 请实现：每1小时执行一次SyncData
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(1 * time.Hour)
	defer tick.Stop()
	stop := make(chan struct{})
	go func() {
		time.Sleep(24 * time.Hour)
		close(stop)
	}()
	for {
		select {
		case <-tick.C:
			err := SyncData()
			if err != nil {
				fmt.Errorf("出错了 %v", err)
			}
		case <-stop:
			return
		}
	}

}

func SyncData() error {
	return nil
}
