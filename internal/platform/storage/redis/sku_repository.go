package redis

type redisSkuRepository struct {
}

func NewRedisSkuRepository() redisSkuRepository {
	return redisSkuRepository{}
}

func (r redisSkuRepository) Save(sku string) error {

	return nil
}
