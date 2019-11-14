package main

import (
	"fmt"
	"net"
	"time"
)

func main(){

	//data := make(map[string]interface{})
	//data["cmd"] = "connect"
	//data["uuid"] = "123456"
	//
	//tmp,_ := json.Marshal(data)
	//fmt.Println(string(tmp))
	//return

	fmt.Println("The Server will be start")
	//客户连接处理
	go processClient()
	go testPrint()
	err := startServer()

	fmt.Println("Server has error:%s",err.Error())
}

func startServer()error{
	listen,err := net.Listen("tcp","localhost:8090")
	if err != nil{
		//记录错误日志，退出
		return err
	}
	defer listen.Close()
	fmt.Println("The Server has listening...")
	for{
		conn,err := listen.Accept()
		if err != nil{
			continue
		}
		fmt.Println("server receive a client connection!")
		go handleConnection(conn)
	}
	return nil
}

func testPrint(){
	for{
		fmt.Printf("当前客户池数量:%d\n",len(clientPool))
		time.Sleep(5*time.Second)
	}


}
