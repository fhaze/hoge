package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	srv := NewService()
	res, err := srv.Get()

	assert.NoError(t, err)
	assert.Equal(t, "msg from mock", res)
}
