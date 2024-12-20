package asa

import (
	"fmt"
	"testing"
)

// go test -v -run TestGetAllCampaigns
func TestGetAllCampaigns(t *testing.T) {
	t.Parallel()

	auth, err := NewTokenConfig(clientID, teamID, keyID, privateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}
	auth.SetOrgID(8276820)
	res, err := NewClientWithAuth(auth).Campaigns.GetAllCampaigns(&GetAllCampaignQuery{
		Limit:  2,
		Offset: 0,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(res.Campaigns)
}

// go test -v -run TestFindCampaigns
func TestFindCampaigns(t *testing.T) {
	t.Parallel()

	auth, err := NewTokenConfig(clientID, teamID, keyID, privateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}
	auth.SetOrgID(0)
	res, err := NewClientWithAuth(auth).Campaigns.FindCampaigns(&Selector{
		Conditions: []*Condition{
			{
				Field:    "name",
				Operator: "IN",
				Values:   []string{""},
			},
		},
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(res.Campaigns)
}
