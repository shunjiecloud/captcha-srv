package modules

import (
	"fmt"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2/config"
	"github.com/shunjiecloud/captcha-srv/store"
)

type moduleWrapper struct {
	Redis *redis.Client
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  redis
	var rConfig RedisConfig
	if err := config.Get("config", "redis").Scan(&rConfig); err != nil {
		panic(err)
	}
	ModuleContext.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", rConfig.Address, rConfig.Port),
		Password: "", // no password set
		DB:       1,  // redis中统计keys数量，需要用dbsize。keys和scan都不太好o(n)。因此，让captcha独占一个db吧。
	})
	_, err := ModuleContext.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	//  captcha
	var captchaConfig CaptchaConfig
	if err := config.Get("config", "captcha").Scan(&captchaConfig); err != nil {
		panic(err)
	}
	store := store.NewRedisCaptchaStore(ModuleContext.Redis, captchaConfig.MaxCollectNum, time.Duration(10)*time.Minute)
	captcha.SetCustomStore(store)
}
