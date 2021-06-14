package sku

import (
	"context"
	"errors"
	"github.com/raulquiros/tcp_server/kit/query"
)

type ListSkuQueryHandler struct {
	skuRepository SkuRepository
}

func NewListSkuQueryHandler(skuRepository SkuRepository) ListSkuQueryHandler {
	return ListSkuQueryHandler{
		skuRepository: skuRepository,
	}
}

func (h ListSkuQueryHandler) Handle(ctx context.Context, query query.Query) error {
	_, ok := query.(ListSkuQuery)
	if !ok {
		return errors.New("unexpected command")
	}
	_, err := h.skuRepository.FindAll(ctx)
	if err != nil {
		return err
	}

	return nil
}
