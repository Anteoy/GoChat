package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type Redis struct {
	client *redis.Conn
}

func (r *Redis) Incr(name) int {
	result, err := r.client.Do(INCR, name)
	if err != nil {
		log.Panic(err)
	}
	return result
}

func (r *Redis) Close() {
	r.client.Close()
}
