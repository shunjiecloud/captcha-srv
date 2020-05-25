package modules

import (
	"fmt"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2/config"
)

type moduleWrapper struct {
	Redis *redis.Client
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  redis
	var host Host
	if err := config.Get("hosts", "redis").Scan(&host); err != nil {
		panic(err)
	}
	ModuleContext.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host.Address, host.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := ModuleContext.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	//  captcha
	store := NewRedisCaptchaStore(10000, time.Duration(10)*time.Minute)
	captcha.SetCustomStore(store)
}
