package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/raulquiros/tcp_server/internal/sku"
)

const SkuRepositoryKey = "skus"

type RedisSkuRepository struct {
	conn *redis.Client
}

func NewRedisSkuRepository(conn *redis.Client) RedisSkuRepository {
	return RedisSkuRepository{conn: conn}
}

func (r RedisSkuRepository) Save(ctx context.Context, sku sku.Sku) error {
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
