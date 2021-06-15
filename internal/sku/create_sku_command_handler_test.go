package sku

import (
	"context"
	"errors"
	"github.com/raulquiros/tcp_server/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_Create_Sku_Command_Handler(t *testing.T) {
	skuRepository := new(storagemocks.SkuRepository)

	t.Run("given an invalid sku it returns the expected response", func(t *testing.T) {
		cmd := givenACreateSkuCommandWithInvalidSkuValue(skuRepository)
		err := whenHandleCommand(skuRepository, cmd)

		assert.Nil(t, err)
	})

	t.Run("given an valid sku it returns the expected response", func(t *testing.T) {
		cmd := givenACreateSkuCommandWithValidSkuValue(skuRepository)
		err := whenHandleCommand(skuRepository, cmd)

		assert.Nil(t, err)
	})

	t.Run("given a sku it returns an error", func(t *testing.T) {
		cmd := givenACreateSkuCommandWithAnError(skuRepository)
		err := whenHandleCommand(skuRepository, cmd)

		assert.Error(t, err)
	})
}

func whenHandleCommand(skuRepository *storagemocks.SkuRepository, cmd CreateSkuCommand) error {
	handler := NewCreateSkuCommandHandler(skuRepository)
	err := handler.Handle(context.Background(), cmd)

	return err
}

func givenACreateSkuCommandWithInvalidSkuValue(skuRepository *storagemocks.SkuRepository) CreateSkuCommand {
	skuRepository.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(nil).Once()

	return NewCreateSkuCommand("test")
}

func givenACreateSkuCommandWithValidSkuValue(skuRepository *storagemocks.SkuRepository) CreateSkuCommand {
	skuRepository.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(nil).Once()

	return NewCreateSkuCommand("AAAA-1111")
}

func givenACreateSkuCommandWithAnError(skuRepository *storagemocks.SkuRepository) CreateSkuCommand {
	skuRepository.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(errors.New("connection refused")).Once()

	return NewCreateSkuCommand("AAAA-1111")
}
