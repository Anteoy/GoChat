package mynet

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
	"time"
)

var (
	RedisPool *redis.Pool
	server    = "139.129.4.187:6379"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	paras := r.URL.Query()
	mingl := paras["m"][0]

	client := RedisPool.Get()
	defer client.Close()

	v, _ := client.Do("GET", mingl)
	fmt.Printf("hello %s\n", v)
	fmt.Fprintf(w, "Hello! %s", v)
}

func InitPool() {
	RedisPool = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10000,
		IdleTimeout: time.Duration(5) * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				log.Printf("Dail redis server %s %v", server, err)
				return nil, err
			}
			if _, err := c.Do("AUTH", "123456"); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("PING"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}
