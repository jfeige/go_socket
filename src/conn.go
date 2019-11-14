package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()
	fmt.Println("server begin process client connection!")
	tmpBuffer := make([]byte,0)
	readerChannel := make(chan []byte,16)
	go readerMessage(readerChannel,conn)

	buffer := make([]byte,1024)
	for{
		n,err := conn.Read(buffer)
		if err != nil{
			if err == io.EOF{
				//客户端主动断开连接
				fmt.Println("客户端主动断开连接")
				return
			}else{
				fmt.Printf("find error:%s",err.Error())
				return
			}
		}
		tmpBuffer = unPack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}


//uuid string					//客户唯一标识
//lastCmd string 				//客户最近一次执行命令
//status int					//在线状态 1:在线;0:离线
//connectTime int				//连接时间
//clientIp	string			//连接ip
//activeTime int64			//最后活跃时间
//lastHeartTime int64			//最后一次心跳时间
//heartCnt	int				//心跳次数
//lastMsgTime int64			//最后一次消息时间
//unConnectTime int64 		//客户主动断开连接时间

/**
	处理接收到的数据
 */
func readerMessage(readerChannel chan []byte,conn net.Conn){
	for{
		select {
		case msg :=<- readerChannel:
			data := make(map[string]interface{})
			err := json.Unmarshal(msg,&data)
			fmt.Println("-----",data)
			if err != nil{
				continue
			}
			uuid := data["uuid"].(string)
			cmd := data["cmd"].(string)
			ip := conn.RemoteAddr()
			content := uuid + "," + cmd + "," + ip.String()
			clientChannel <- content
		}
	}
}