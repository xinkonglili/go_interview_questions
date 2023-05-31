/*使用go实现一个简单的本地消息队列// 以下为伪代码
interface LocalMQ {
    Send(...) error
    StartUp(...) error
    RegisterListener(...) error
}*/

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	ID   int
	Data string
}

type LocalMQ interface {
	Send(msg *Message) error
	StartUp()
	RegisterListener(listener func(msg *Message) error)
	ShutDown()
}

type InMemoryMQ struct {
	mutex       sync.Mutex
	listeners   []func(msg *Message) error
	messageChan chan *Message
}

func (mq *InMemoryMQ) Send(msg *Message) error {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.messageChan <- msg
	return nil
}

func (mq *InMemoryMQ) StartUp() {
	go func() {
		for {
			select {
			case msg := <-mq.messageChan:
				for _, listener := range mq.listeners {
					err := listener(msg)
					if err != nil {
						log.Printf("监听器错误: %v", err)
					}
				}
			default:
				if mq.messageChan == nil {
					return
				}
			}
		}
	}()
}

func (mq *InMemoryMQ) RegisterListener(listener func(msg *Message) error) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.listeners = append(mq.listeners, listener)
}

func (mq *InMemoryMQ) ShutDown() {
	close(mq.messageChan)
}
func (mq *InMemoryMQ) init() {
	mq.messageChan = make(chan *Message)
}

func main() {
	mq := &InMemoryMQ{}
	//确保StartUp的时候，通道已经被初始化，否则会出现读取的通道为nil
	mq.init()
	mq.RegisterListener(func(msg *Message) error {
		fmt.Printf("消息体Id：%v, Data：%v\n", msg.ID, msg.Data)
		return nil
	})
	mq.StartUp()
	msg := &Message{
		ID:   11,
		Data: "Go language",
	}
	mq.Send(msg)
	time.Sleep(time.Second)
	mq.ShutDown()
}
