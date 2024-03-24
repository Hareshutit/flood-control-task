package repository

import (
	"github.com/gomodule/redigo/redis"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis_cart:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func CreateRedisRepository() RedisRepository {
	pool := newPool()
	connect := pool.Get()

	return RedisRepository{connect}
}
