package redis

import (
	repo "github.com/baxromumarov/my-services/api-gateway/storage/repo"
	redis "github.com/gomodule/redigo/redis"
)

type redisRepo struct {
	rConn *redis.Pool
}

func NewRedisRepo(rds *redis.Pool) repo.RedisRepositoryStorage {
	return &redisRepo{
		rConn: rds,
	}
}

// Set
func (r *redisRepo) Set(key, value string) error {
	conn := r.rConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

// SetWithTTL

func (r *redisRepo) SetWithTTL(key, value string, second int64) error {
	conn := r.rConn.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, second, value)
	return err
}

// Get
func (r *redisRepo) Get(key string) (interface{}, error) {
	conn := r.rConn.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}
