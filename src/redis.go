package main

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var(
	redis_host string = "123.56.85.167"
	redis_port int = 6379
	redis_passwd string = ""
)
func initRedisPool()*redis.Pool{
	redis_address := redis_host + ":" + strconv.Itoa(redis_port)
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn,err := redis.Dial("tcp",redis_address)
			if err != nil{
				return nil,err
			}
			if redis_passwd != ""{
				err = conn.Send("AUTH",redis_passwd)
				if err != nil{
					return nil,err
				}
			}
			return conn,nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			return c.Send("PING")
		},
	}

	return pool
}