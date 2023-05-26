package utils_kratos_auth_match

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}

func TestExample(t *testing.T) {
	var e error
	if !assert.NoError(t, e) {
		return
	}
	var b = true
	if !assert.True(t, b) {
		return
	}
}
