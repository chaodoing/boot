package captcha

import (
	"log"
	"strings"
	"time"
	
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type Store struct {
	// Duration 验证码有效时间
	Duration time.Duration `json:"duration"`
	rdx      *redis.Client
	prefix   string
}

func (s Store) Set(idKey string, value string) (err error) {
	idKey = s.prefix + ":" + idKey
	_, err = s.rdx.Set(idKey, value, s.Duration*time.Minute).Result()
	return
}

func (s Store) Get(idKey string, clear bool) (value string) {
	var err error
	value, err = s.rdx.Get(idKey).Result()
	if err != nil {
		log.Println(err)
	}
	if clear {
		_, err := s.rdx.Del(idKey).Result()
		if err != nil {
			log.Println(err)
		}
		return
	}
	return
}

func (s Store) Verify(idKey, value string, clear bool) bool {
	idKey = s.prefix + ":" + idKey
	var verifyCode = s.Get(idKey, clear)
	return strings.EqualFold(verifyCode, value)
}

func NewStore(rdx *redis.Client, duration int64, prefix ...string) base64Captcha.Store {
	var index string
	if len(prefix) > 0 {
		index = prefix[0]
	} else {
		index = "captcha"
	}
	return &Store{
		Duration: time.Duration(duration),
		rdx:      rdx,
		prefix:   index,
	}
}
