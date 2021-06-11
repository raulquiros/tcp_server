package sku

import (
	"context"
	"github.com/raulquiros/tcp_server/kit/command"
)

type createSkuCommandHandler struct {
	skuRepository SkuRepository
}

func NewCreateSkuCommandHandler(skuRepository SkuRepository) createSkuCommandHandler {
	return createSkuCommandHandler{
		skuRepository: skuRepository,
	}
}

func (h createSkuCommandHandler) Handle(context context.Context, cmd command.Command) error {
	return nil
}
