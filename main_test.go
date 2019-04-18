package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainExec(t *testing.T) {
	assert.NotPanics(t, main)
}
