package float

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstNotNilTest(t *testing.T) {
	var a *float64 = nil
	b := New(2.3)
	result := FirstNotNil(a, b)
	assert.Equal(t, 2.3, *result)

	var c *float64 = nil
	result2 := FirstNotNil(a, c)
	assert.Nil(t, result2)
}
