package pointer

import (
	"github.com/gozelle/lancet/internal"
	"testing"
)

func TestExtractPointer(t *testing.T) {
	
	assert := internal.NewAssert(t, "TestExtractPointer")
	
	a := 1
	b := &a
	c := &b
	d := &c
	
	assert.Equal(1, ExtractPointer(d))
}
