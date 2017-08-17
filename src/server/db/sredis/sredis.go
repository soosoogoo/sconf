package sredis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type RedisDriver struct {
	Database string
	Conn     redis.Conn
}

func (rd RedisDriver) Connent(databae string) (redis.Conn, error) {

	//redis.DialPassword("")加密码
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	rd.Conn = c
	rd.Database = databae
	return c, err
}

func (rd RedisDriver) Hset(name string, key interface{}, value string) {
	_, err := rd.Conn.Do("HSET", name, key, value)

	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func (rd RedisDriver) Hget(name string, key interface{}) {
	a, err := rd.Conn.Do("HGET", name, key)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	fmt.Println(a)
}

// 写入值永不过期
//	_, err = c.Do("HSET", "username", "1", "nick")
//	if err != nil {
//		fmt.Println("redis set failed:", err)
//	}
//	username, err := redis.String(c.Do("HGET", "username", "1"))
//	if err != nil {
//		fmt.Println("redis get failed:", err)
//	} else {
//		fmt.Printf("Got username %v \n", username)
//	}