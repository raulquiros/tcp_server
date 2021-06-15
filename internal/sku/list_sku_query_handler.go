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

func (h ListSkuQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	_, ok := q.(ListSkuQuery)
	if !ok {
		return query.Response{}, errors.New("unexpected command")
	}

	var valid, invalid, duplicated int

	skuCollection, err := h.skuRepository.FindAll(ctx)
	if err != nil {
		return query.Response{}, err
	}

	for _, sku := range skuCollection {

		skuObj, err := NewSKU(sku)
		if err != nil {
			invalid++
		}

		if isDuplicated(skuObj.value) {
			duplicated++
		}

		valid++
	}

	return query.Response{
		Data: map[string]int{
			"valid":      valid,
			"invalid":    invalid,
			"duplicated": duplicated,
		},
	}, nil
}

var duplicates = make(map[string]bool)

func isDuplicated(sku string) bool {

	_, exist := duplicates[sku]

	if exist {
		return true
	}

	duplicates[sku] = true

	return false
}
