package asa

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -run TestCreateAdGroup
func TestCreateAdGroup(t *testing.T) {
	t.Parallel()

	client, err := initClient()
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	data := new(AdGroup)
	data.Name = "testAdGroup123"
	data.Status = AdGroupStatusEnabled
	data.StartTime = DateTime{time.Now()}

	res, err := client.AdGroups.CreateAdGroup(0, data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(res.AdGroup)
}
