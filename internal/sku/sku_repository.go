package sku

type SkuRepository interface {
	Save(sku string) error
}
