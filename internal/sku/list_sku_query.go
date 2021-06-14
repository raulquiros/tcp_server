package sku

import (
	"github.com/raulquiros/tcp_server/kit/query"
)

const SkuListType query.Type = "query.sku.list"

type ListSkuQuery struct{}

func (c ListSkuQuery) Type() query.Type {
	return SkuListType
}

func NewListSkuQuery() ListSkuQuery {
	return ListSkuQuery{}
}
