package sku

import (
	"github.com/raulquiros/tcp_server/kit/command"
)

const SkuCreateType command.Type = "command.user.create"

type CreateSkuCommand struct {
	sku string
}

func (c CreateSkuCommand) Type() command.Type {
	return SkuCreateType
}

func NewCreateSkuCommand(sku string) CreateSkuCommand {
	return CreateSkuCommand{sku: sku}
}
