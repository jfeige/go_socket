package main

import (
	"github.com/garyburd/redigo/redis"
	"sync"
)


var(
	pool *redis.Pool
	lock     = sync.Mutex{}
	clientPool map[string]*Client
)



func init(){
	//初始化redis连接池
	pool = initRedisPool()
	//初始化客户连接池
	clientPool = make(map[string]*Client)

}








func  GetRedisConn()redis.Conn{
	return pool.Get()
}