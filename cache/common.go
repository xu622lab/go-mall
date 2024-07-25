/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-25 15:29:49
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 15:53:22
 * @FilePath: /go-mall/cache/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("redis config err", err)
	}
	LoadRedisData(file)
	Redis()
}

func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		// Password: ,
		DB: int(db),
	})
	_, err := client.Ping().Result() // 心跳检测
	if err != nil {
		panic(err)
	}

	RedisClient = client
}
