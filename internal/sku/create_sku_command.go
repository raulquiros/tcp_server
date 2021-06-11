package sku

import (
	"github.com/raulquiros/tcp_server/kit/command"
)

const SkuCreateType command.Type = "command.user.create"

type createSkuCommand struct {
	sku string
}

func (o createSkuCommand) Type() command.Type {
	return SkuCreateType
}

func NewCreateSkuCommand(sku string) createSkuCommand {
	return createSkuCommand{sku: sku}
}
