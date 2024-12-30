package asa

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func initClient() (*Client, error) {
	client := NewClient(nil, os.Getenv("AppleAccessToken"))
	orgID, err := strconv.ParseInt(os.Getenv("AppleOrgID"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse orgID: %v", err)
	}
	err = client.SetOrgID(orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to set orgID: %v", err)
	}
	err = client.SetHTTPDebug(true)
	if err != nil {
		return nil, fmt.Errorf("failed to set http debug: %v", err)
	}
	return client, nil
}

// go test -v -run TestSearchApps
func TestSearchApps(t *testing.T) {
	t.Parallel()

	client, err := initClient()
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	res, err := client.App.SearchApps(&SearchAppsQuery{
		Query:           "run",
		Limit:           10,
		Offset:          0,
		ReturnOwnedApps: true,
	})
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	fmt.Printf("%#v\n", res.AppInfos)
}
