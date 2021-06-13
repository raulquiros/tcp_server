package sku

import (
	"context"
	"errors"
	"github.com/raulquiros/tcp_server/kit/command"
)

type CreateSkuCommandHandler struct {
	skuRepository SkuRepository
}

func NewCreateSkuCommandHandler(skuRepository SkuRepository) CreateSkuCommandHandler {
	return CreateSkuCommandHandler{
		skuRepository: skuRepository,
	}
}

func (h CreateSkuCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createSkuCommand, ok := cmd.(CreateSkuCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.skuRepository.Save(ctx, createSkuCommand.sku)
}
