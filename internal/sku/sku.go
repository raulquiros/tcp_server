package sku

import (
	"fmt"
	"regexp"
)

type Sku struct {
	value string
}

func NewSKU(value string) (Sku, error) {

	match, err := regexp.Match("^[A-Z]{4}-[0-9]{4}$", []byte(value))
	if err != nil || !match {
		return Sku{}, fmt.Errorf("Invalid SKU value: %s", value)
	}

	return Sku{value: value}, nil
}
