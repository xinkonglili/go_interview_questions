/*
func DoLoadAll() error
func DoSyncData() error
// 请实现以下接口，

	type Syncer interface {
		Notify()
		Startup()
	}

// 用以支撑如下能力
syncer := &XXXSyncer()
1)调用 syncer.Startup() 的时候，会调用 DoLoadAll(), 并且保证一个进程内只会调用一次；
2)当调用 syncer.Notify() 时，会调用 DoSyncData(), 并且保证不会并发调用。
*/
package main

import (
	"fmt"
	"sync"
)

type Syncer interface {
	Notify()
	Startup()
}

type XXXSyncer struct {
	syncOnce sync.Once
	mutex    sync.Mutex
}

func (x *XXXSyncer) Startup() {
	x.syncOnce.Do(func() {
		err := DoLoadAll()
		if err != nil {
			fmt.Println("数据加载失败")
		} else {
			fmt.Println("数据加载成功")
		}
	},
	)
}
func (x *XXXSyncer) Notify() {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	err := DoSyncData()
	if err != nil {
		fmt.Println("出现并发调用，数据同步失败")
	} else {
		fmt.Println("数据同步成功")
	}
}

func DoLoadAll() error {
	return nil
}

func DoSyncData() error {
	return nil
}
func main() {
	syncer := &XXXSyncer{}
	syncer.Startup()
	syncer.Notify()
}
