package strconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInt(t *testing.T) {
	var err error
	var s string
	var converted int
	assert := assert.New(t)

	s = "$1,000,000.1"
	converted, err = ConvertInt(s)
	if assert.Nil(err) {
		assert.Equal(1000000, converted)
	}

	s = "$1,000,000.5"
	converted, err = ConvertInt(s)
	if assert.Nil(err) {
		assert.Equal(1000001, converted)
	}

	s = "1.00.00"
	converted, err = ConvertInt(s)
	assert.NotNil(err)
}

func TestConvertInt64(t *testing.T) {
	var err error
	var s string
	var converted int64
	assert := assert.New(t)

	s = "$1,000,000.00"
	converted, err = ConvertInt64(s)
	if assert.Nil(err) {
		assert.Equal(int64(1000000), converted)
	}
}

func TestConvertFloat64(t *testing.T) {
	var err error
	var s string
	var converted float64
	assert := assert.New(t)

	s = "$1,000,000.00"
	converted, err = ConvertFloat64(s)
	if assert.Nil(err) {
		assert.Equal(1000000.00, converted)
	}
}
