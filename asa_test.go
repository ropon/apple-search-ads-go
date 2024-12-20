package asa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test -v -run TestNewClient
func TestNewClient(t *testing.T) {
	t.Parallel()

	c := NewClient(nil)

	assert.Equal(t, defaultBaseURL, c.client.BaseUrl.String())

	assert.False(t, c.client.Debug)
	c.SetHTTPDebug(true)
	assert.True(t, c.client.Debug)

}
