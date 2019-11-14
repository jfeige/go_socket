package main

import "time"




/**
	处理心跳数据
 */
func handleHeart(uuid string){

	client,ok := getClient(uuid)
	if !ok{
		//客户池中不存在，则重新加入客户池
	}
	client.heartCnt += 1
	client.lastHeartTime = time.Now().Unix()
	client.activeTime = time.Now().Unix()
	addClient(client)
}
