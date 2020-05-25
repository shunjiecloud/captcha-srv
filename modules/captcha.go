package modules

import (
	"time"
)

//RedisCaptchaStorePrefix 验证码种子集合前缀
const RedisCaptchaStorePrefix = "captcha:collects:"

//RedisCaptchaMaxNumKey 保有的验证码最大数量
const RedisCaptchaMaxNumKey = "captcha:max"

//RedisCaptchaCurNumKey 当前验证码数量
const RedisCaptchaCurNumKey = "captcha:cur"

type RedisCaptchaStore struct {
	collectNum int
	expiration time.Duration
}

func (store *RedisCaptchaStore) Set(id string, digits []byte) {
	//  TODO ： 检查key数量是否超过上限
	key := RedisCaptchaStorePrefix + id
	ModuleContext.Redis.Set(key, digits, store.expiration)
}

func (store *RedisCaptchaStore) Get(id string, clear bool) (digits []byte) {
	digits = make([]byte, 0)
	key := RedisCaptchaStorePrefix + id
	ret, err := ModuleContext.Redis.Get(key).Result()
	if err != nil {
		return nil
	}
	digits = []byte(ret)
	return digits
}

func NewRedisCaptchaStore(collectNum int, expiration time.Duration) *RedisCaptchaStore {
	store := RedisCaptchaStore{
		collectNum: collectNum,
		expiration: expiration,
	}
	return &store
}
