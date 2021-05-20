package persistence

import (
	"opencodes/tools"
	"time"

	"github.com/go-redis/redis"
)

// RedisStore represents the cache with redis persistence
type RedisStore struct {
	pool              *redis.Client
	defaultExpiration time.Duration
}

// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(host string, password string, defaultExpiration time.Duration) *RedisStore {
	pool := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	return &RedisStore{pool, defaultExpiration}
}

// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCacheWithPool(pool *redis.Client, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{pool, defaultExpiration}
}

// Set (see CacheStore interface)
func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {

	return c.pool.Set(key, value, expires).Err()
}

// Add (see CacheStore interface)
func (c *RedisStore) Add(key string, value interface{}) error {

	return c.pool.Append(key, value.(string)).Err()
}

// Replace (see CacheStore interface)
func (c *RedisStore) Replace(key string, value interface{}, expires time.Duration) error {
	return c.pool.Set(key, value, expires).Err()

}

// Get (see CacheStore interface)
func (c *RedisStore) Get(key string, ptrValue interface{}) error {

	raw, err := c.pool.Get(key).Result()
	if raw == "" {
		return ErrCacheMiss
	}

	if err != nil {
		return err
	}
	return tools.Deserialize([]byte(raw), ptrValue)
}

func exists(conn *redis.Client, key string) bool {
	retval, _ := conn.Exists(key).Result()
	return retval == 0
}

// Delete (see CacheStore interface)
func (c *RedisStore) Delete(key string) error {

	return c.pool.Del(key).Err()
}

// Increment (see CacheStore interface)
func (c *RedisStore) Increment(key string, delta int64) (int64, error) {

	return c.pool.IncrBy(key, delta).Result()
}

// Decrement (see CacheStore interface)
func (c *RedisStore) Decrement(key string, delta int64) (newValue int64, err error) {

	return c.pool.DecrBy(key, delta).Result()
}

// Flush (see CacheStore interface)
func (c *RedisStore) Flush() error {

	defer c.pool.Close()
	_, err := c.pool.FlushAll().Result()
	return err
}

//func (c *RedisStore) invoke(f func(string, ...interface{}) (interface{}, error),
//	key string, value interface{}, expires time.Duration) error {
//
//	switch expires {
//	case DEFAULT:
//		expires = c.defaultExpiration
//	case FOREVER:
//		expires = time.Duration(0)
//	}
//
//	b, err := tools.Serialize(value)
//	if err != nil {
//		return err
//	}
//
//	if expires > 0 {
//		_, err := f("SETEX", key, int32(expires/time.Second), b)
//		return err
//	}
//
//	_, err = f("SET", key, b)
//	return err
//
//}
