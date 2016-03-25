package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"mynet"
	"strconv"
)

func Incr(name string) int {
	c := RedisPool.Get()
	defer c.Close()

	result, err := redis.Int(c.Do("INCR", name))
	if err != nil {
		log.Panic(err)
	}
	return result
}

func Get(name string) string {
	c := RedisPool.Get()
	defer c.Close()

	return GetFromConn(c, name)
}

func GetFromConn(c redis.Conn, name string) string {

	result, err := redis.String(c.Do("GET", name))
	if err != nil {
		log.Panicln("redis GET " + name + " is error")
	}

	return result
}

func Set(name string, value string) int {
	c := RedisPool.Get()
	defer c.Close()

	return SetFromConn(c, name, value)
}

func SetFromConn(c redis.Conn, name string, value string) int {

	_, err := c.Do("SET", name, value)
	if err != nil {
		return 0
	}
	return 1
}

func AddUser(user *mynet.User) bool {
	c := RedisPool.Get()
	defer c.Close()

	id := Incr("NO")
	count := 0
	count += SetFromConn(c, "user:"+strconv.Itoa(id)+":name", user.Name)
	count += SetFromConn(c, "user:"+strconv.Itoa(id)+":pass", user.Pass)
	count += SetFromConn(c, "user:"+strconv.Itoa(id)+":friends", user.Friends)
	count += SetFromConn(c, "user:"+strconv.Itoa(id)+":other", user.Other)
	if count != 4 {
		return false
	}
	return true
}

func GetUser(id int) *mynet.User {
	c := RedisPool.Get()
	defer c.Close()

	name := GetFromConn(c, "user:"+strconv.Itoa(id)+":name")
	pass := GetFromConn(c, "user:"+strconv.Itoa(id)+":pass")
	friends := GetFromConn(c, "user:"+strconv.Itoa(id)+":friends")
	other := GetFromConn(c, "user:"+strconv.Itoa(id)+":other")
	defer func() {
		if r := recover(); r != nil {
			log.Panicln("获取id为" + strconv.Itoa(id) + "用户信息出错")
			log.Panic(r)
		}
	}()
	user := mynet.NewUser(id, name, pass, friends, other)
	return user
}
