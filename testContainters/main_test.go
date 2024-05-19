package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
    t.Log(("We are testing 'testing containers' package."))
    assert.Equal(t, 1, 1, "The two numbers should be the same.")
}
