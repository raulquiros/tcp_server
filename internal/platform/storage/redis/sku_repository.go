package redis

import "github.com/go-redis/redis/v8"

type redisSkuRepository struct {
	conn *redis.Client
}

func NewRedisSkuRepository(conn *redis.Client) redisSkuRepository {
	return redisSkuRepository{conn: conn}
}

func (r redisSkuRepository) Save(sku string) error {

	return nil
}
