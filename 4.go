/*
请帮助提供一个生成ID的接口。
约束：有多台服务器；QPS每秒: 100000

	interface IDFactory {
	    GetID() (string, error)
	}
*/
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
)

const (
	idLength = 16
)

type IDFactory interface {
	GetID() string
}
type RandomIDFactory struct {
	mutex sync.Mutex
}

// 生成一个新的唯一ID
func (f *RandomIDFactory) GetID() string {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	randomBytes := make([]byte, idLength)
	rand.Read(randomBytes)
	idString := base64.URLEncoding.EncodeToString(randomBytes)
	return idString
}

func main() {
	// 1、创建
	factory := &RandomIDFactory{}
	// 2、生成
	id := factory.GetID()
	fmt.Println(id)
}
