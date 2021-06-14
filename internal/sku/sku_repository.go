package sku

import "context"

type SkuRepository interface {
	FindAll(ctx context.Context) ([]string, error)
	Save(ctx context.Context, sku string) error
}
