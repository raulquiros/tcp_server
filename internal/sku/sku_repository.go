package sku

import "context"

type SkuRepository interface {
	FindAll(ctx context.Context) ([]string, error)
	Save(ctx context.Context, sku string) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/storagemocks --name=SkuRepository
