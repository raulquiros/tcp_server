package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

const SkuRepositoryKey = "skus"

type RedisSkuRepository struct {
	conn *redis.Client
}

func NewRedisSkuRepository(conn *redis.Client) RedisSkuRepository {
	return RedisSkuRepository{conn: conn}
}

func (r RedisSkuRepository) FindAll(ctx context.Context) ([]string, error) {
	resp := r.conn.LRange(ctx, SkuRepositoryKey, 0, -1)

	return resp.Result()
}

func (r RedisSkuRepository) Save(ctx context.Context, sku string) error {
	p, err := json.Marshal(sku)
	if err != nil {
		return err
	}

	resp := r.conn.LPush(ctx, SkuRepositoryKey, p)
	if resp.Err() != nil {
		return resp.Err()
	}

	return nil
}
