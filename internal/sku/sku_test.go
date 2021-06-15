package sku

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSku(t *testing.T) {
	t.Run("given an invalid sku it returns an error", func(t *testing.T) {
		_, err := NewSKU("test")

		assert.Error(t, err)
	})

	t.Run("given a valid sku it returns a sku obj", func(t *testing.T) {
		obj, err := NewSKU("AAAA-1111")

		assert.Nil(t, err)
		assert.Equal(t, "AAAA-1111", obj.value)
	})
}
