package main

import (
	"fmt"
	"strings"
	"time"
)

var(
	clientChannel = make(chan string,1000)
)

type Client struct {
	uuid string					//客户唯一标识
	lastCmd string 				//客户最近一次执行命令
	status int					//在线状态 1:在线;0:离线
	connectTime int64				//连接时间
	clientIp	string			//连接ip
	activeTime int64			//最后活跃时间
	lastHeartTime int64			//最后一次心跳时间
	heartCnt	int				//心跳次数
	lastMsgTime int64			//最后一次消息时间
	unConnectTime int64 		//客户主动断开连接时间
}

/**
#lf#000001{"cmd":"connect","uuid":"123456"}
#lf#000001{"cmd":"msg","uuid":"123456","content":"消息内容","receiveUuid":"888888"}
#lf#000001{"cmd":"disconnect","uuid":"123456"}
#lf#000001{"cmd":"heart","uuid":"123456"}
 */

func processClient(){
	for{
		select {
		case content :=<- clientChannel:
			data := strings.Split(content,",")
			uuid := data[0]
			lastCmd := data[1]
			if lastCmd == "connect"{
				//新客户连接
				ip := data[2]
				client := new(Client)
				client.uuid = uuid
				client.lastCmd = lastCmd
				client.status = 1
				client.connectTime = time.Now().Unix()
				client.clientIp = ip
				client.activeTime = time.Now().Unix()
				addClient(client)
			}else{
				tmpClient,ok := getClient(uuid)
				if ok{
					tmpClient.lastCmd = lastCmd
					tmpClient.activeTime = time.Now().Unix()
					tmpClient.status = 1
				}
				if lastCmd == "msg"{
					//发送消息
					tmpClient.lastMsgTime = time.Now().Unix()
				}else if lastCmd == "disconnect"{
					//客户端主动断开连接
					tmpClient.unConnectTime = time.Now().Unix()
					tmpClient.status = 0
				}else if lastCmd == "heart"{
					//心跳
					tmpClient.heartCnt = tmpClient.heartCnt + 1
					tmpClient.lastHeartTime = time.Now().Unix()
				}
				addClient(tmpClient)
			}
		case <- time.After(time.Second * 5):
			//5秒清理一次客户池
			fmt.Println("开始清理客户池....")
			go removeOffLineClient()
		}
	}
}


/**
	新连接客户端加入客户池
 */
func addClient(client *Client){
	defer lock.Unlock()
	lock.Lock()
	clientPool[client.uuid] = client
}

/**
	根据客户唯一标识获取客户对象
 */
func getClient(uuid string) (*Client,bool){
	defer lock.Unlock()
	lock.Lock()
	client,ok := clientPool[uuid]
	return client,ok
}

/**
	更新客户端信息
 */
func updateClient(){
	//加锁
}

/**
	从客户池中移除
 */
func removeClient(uuid string){
	//加锁
	defer lock.Unlock()
	lock.Lock()
	delete(clientPool,uuid)
}

/**
	获取当前所有在线状态客户列表
 */
func getOnLineClientList()[]*Client{
	defer lock.Unlock()
	lock.Lock()
	clientList := make([]*Client,0)
	for _,client := range clientPool{
		if client.status == 1{
			clientList = append(clientList,client)
		}
	}
	return clientList
}

/**
	定时任务，从客户池中移除离线状态客户
 */
func removeOffLineClient(){
	defer lock.Unlock()
	lock.Lock()
	for _,client := range clientPool{
		if client.status == 0 || (client.activeTime < time.Now().Unix() - 10){
			delete(clientPool,client.uuid)
		}
	}
}