package store

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/shunjiecloud/errors"
	"github.com/shunjiecloud/pkg/log"
)

//RedisCaptchaStorePrefix 验证码集合前缀
const RedisCaptchaStorePrefix = "captcha:collects:"

type RedisCaptchaStore struct {
	maxCollectNum uint64
	expiration    time.Duration
	redis         *redis.Client
}

func (store *RedisCaptchaStore) Set(id string, digits []byte) {
	//  检查key数量是否超过上限
	curNum, err := store.redis.DBSize().Uint64()
	if err != nil {
		log.Error(errors.New("get captcha cur num failed"))
		return
	}
	if curNum >= store.maxCollectNum {
		//  超数，服务不可用
		log.Error(errors.New("captcha num over than max"))
		return
	}
	//  设置key
	key := RedisCaptchaStorePrefix + id
	_, err = store.redis.Set(key, digits, store.expiration).Result()
	if err != nil {
		log.Error(errors.New("captcha store failed"))
		return
	}
}

func (store *RedisCaptchaStore) Get(id string, clear bool) (digits []byte) {
	digits = make([]byte, 0)
	key := RedisCaptchaStorePrefix + id
	ret, err := store.redis.Get(key).Result()
	if err != nil {
		return nil
	}
	digits = []byte(ret)

	if clear == true {
		//  clear为true，删除id
		_, err = store.redis.Del(key).Result()
		if err != nil {
			log.Error(errors.New("delete captcha store failed"))
			return
		}
	}
	return
}

func NewRedisCaptchaStore(redis *redis.Client, maxCollectNum uint64, expiration time.Duration) *RedisCaptchaStore {
	store := RedisCaptchaStore{
		maxCollectNum: maxCollectNum,
		expiration:    expiration,
		redis:         redis,
	}
	return &store
}
