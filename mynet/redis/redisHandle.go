package redis

import (
	redigo "github.com/garyburd/redigo/redis"
	"log"
	"mynet/model/User"
)

type Redis redigo.Conn

func (r *Redis) Incr(name) int {
	result, err := r.Do("INCR", name)
	if err != nil {
		log.Panic(err)
	}
	return result
}

func (r *Redis) Get(name) string {
	result, err := r.Do("GET", name)
	if err != nil {
		log.Panicln("redis GET " + name + " is error")
	}
	return result
}

func (r *Redis) Set(name string, value string) insertNum int {
	_ , err := r.Do("SET", name, value)
	if err != nil {
		insertNum=0
	}
	insertNum=1
}

func (r *Redis) AddUser(user User) bool {
	id := r.Incr()
	_,err:=r.Set("user"+id+"name", user.name)
	_,err:=r.Set("user"+id+"pass", user.pass)
	_,err:=r.Set("user"+id+"friends", user.friends)
	_,err:=r.Set("user"+id+"other", user.other)
	if err!=nil {
		return false
	}
	return true
}

func (r *Redis) GetUser(id int) User {
	name,err   :=r.Get("user"+id+"name", user.name)
	pass,err   :=r.Get("user"+id+"pass", user.pass)
	friends,err:=r.Get("user"+id+"friends", user.friends)
	other,err  :=r.Get("user"+id+"other", user.other)
	 defer func() {  
        if r := recover(); r != nil {  
            log.Panicln("获取id为"+id+"用户信息出错")
            log.Panic(r)
            return nil
        }  
    }()  
	return new(User{id,name,pass,friends,other})
}

func (r *Redis) Close() {
	r.Close()
}
