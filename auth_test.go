package asa

import (
	"fmt"
	"testing"
)

var (
	clientID   = ""
	teamID     = clientID
	keyID      = ""
	privateKey = `

`
)

// go test -v -run TestNewTokenConfig
func TestNewTokenConfig(t *testing.T) {
	t.Parallel()

	auth, err := NewTokenConfig(clientID, teamID, keyID, privateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert.Equal(t, "a", auth.jwtGenerator.clientID)
	clientSecret, err := auth.GenerateClientSecret()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(clientSecret)
}
