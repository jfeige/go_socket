package main

import (
	"fmt"
	"time"
)

var(
	messageChannel chan *Message
)

/**
	消息包
 */
type Message struct {
	sendUuid 		string
	receiveUuid 	string
	sendTime		int64
	content 		string
}



/**
	发送消息
	sendUuid 		消息发送方uuid
	receiveUuid		消息接收方uuid
 */
func sendMsg(sendUuid,receiveUuid,content string){
	msg := new(Message)
	msg.sendUuid = sendUuid
	msg.receiveUuid = receiveUuid
	msg.sendTime = time.Now().Unix()
	msg.content = content
	//加入到消息队列
	if messageChannel == nil{
		messageChannel = make(chan *Message,500)
	}
	messageChannel <- msg
}

/**
	起个单独协程处理消息数据
 */
func handelMessageChan(){
	if messageChannel == nil{
		messageChannel = make(chan *Message,500)
	}
	for {
		msg :=<-messageChannel
		fmt.Println(msg)
	}
}