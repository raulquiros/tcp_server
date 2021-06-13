package sku

import (
	"errors"
	"fmt"
	"regexp"
)

type Sku struct {
	value string
}

func newSKU(value string) (Sku, error) {

	match, err := regexp.Match("^[A-Z]{4}-[0-9]{4}$", []byte(value))
	if err != nil || match == false {
		return Sku{}, errors.New(fmt.Sprintf("Invalid SKU value: %s", value))
	}

	return Sku{value: value}, nil
}
