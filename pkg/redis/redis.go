package redis

import (
	"fmt"
	"opencodes/tools"
	"time"

	goRedis "github.com/go-redis/redis"
)

// 设置
func Set(key string, values interface{}, exporation time.Duration) error {
	cmd := client.Set(key, values, exporation)
	_, err := cmd.Result()
	return err
}

// 设置过期时间
func SetExpire(key string, expire time.Duration) error {
	cmd := client.Expire(key, expire)
	_, err := cmd.Result()
	if err != nil && err != goRedis.Nil {
		return err
	}
	return nil
}

// 获取值
func Get(key string) (string, error) {
	s, err := client.Get(key).Result()
	if err != nil && err != goRedis.Nil {

		return s, err
	}
	fmt.Println(s, "----key")
	err = nil
	return s, err
}

func DecrBy(key string, decrement int64) (int64, error) {
	reply, err := client.DecrBy(key, decrement).Result()
	if err != nil {

	}
	return reply, err
}

func IncrBy(key string, decrement int64) (int64, error) {
	reply, err := client.IncrBy(key, decrement).Result()
	if err != nil {

	}
	return reply, err
}

func HIncrBy(key, field string, decrement int64) (int64, error) {
	r, err := client.HIncrBy(key, field, decrement).Result()
	if err != nil {

	}
	return r, err
}
func Del(key ...string) error {
	_, err := client.Del(key...).Result()

	if err != nil {

	}
	return err
}

func HashSet(key string, value interface{}) error {
	v, _ := tools.StructToMap(value)
	cmd := client.HMSet(key, v)
	_, err := cmd.Result()
	return err
}
func HSet(key, field string, value interface{}) error {
	v, _ := tools.StructToMap(value)
	cmd := client.HSet(key, field, v)
	_, err := cmd.Result()
	return err
}

func HashGetAll(key string) map[string]string {

	cmd := client.HGetAll(key)

	res, _ := cmd.Result()
	return res
}

func PushList(key string, value interface{}) error {
	v, _ := tools.StructToMap(value)
	cmd := client.LPush(key, v)
	_, err := cmd.Result()
	return err
}

func HGet(key, field string) (string, error) {
	cmd := client.HGet(key, field)
	s, err := cmd.Result()
	if err != nil && err != goRedis.Nil {

	}
	return s, err
}
