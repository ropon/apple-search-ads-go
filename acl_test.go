package asa

import (
	"fmt"
	"testing"
)

// go test -v -run TestGetUserACL
func TestGetUserACL(t *testing.T) {
	t.Parallel()

	auth, err := NewTokenConfig(clientID, teamID, keyID, privateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}

	res, err := NewClient(auth).AccessControlList.GetUserACL()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(res.UserAcls[0].OrgName)
}
