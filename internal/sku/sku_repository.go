package sku

import "context"

type SkuRepository interface {
	Save(ctx context.Context, sku string) error
}
