package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceSpace(t *testing.T) {
	var s string
	var converted string
	assert := assert.New(t)

	s = " this  is a   book  "
	converted = ReduceSpace(s)
	assert.Equal("this is a book", converted)

}

func TestRemoveSpace(t *testing.T) {
	var s string
	var converted string
	assert := assert.New(t)

	s = " this  is a   book  "
	converted = RemoveSpace(s)
	assert.Equal("thisisabook", converted)
}
