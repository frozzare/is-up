package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestUp(t *testing.T) {
	assert.Equal(t, true, isUp("google.com"))
}

func TestDown(t *testing.T) {
	assert.Equal(t, false, isUp("hpaghepgapkejkgjkaegköage.com"))
}
