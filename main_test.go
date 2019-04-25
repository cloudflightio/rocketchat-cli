package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMainExec(t *testing.T) {
	assert.NotPanics(t, main)
}

func TestMainPanic(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"foo", "%qfoo=bar"}

	assert.Panics(t, main)
}
