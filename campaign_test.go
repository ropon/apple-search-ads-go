package asa

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -run TestGetAllCampaigns
func TestGetAllCampaigns(t *testing.T) {
	t.Parallel()

	auth, err := NewTokenConfig(clientID, teamID, keyID, privateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}
	res, err := NewClient(auth).Campaigns.GetAllCampaigns(&GetAllCampaignQuery{
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
	res, err := NewClient(auth).Campaigns.FindCampaigns(&Selector{
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

// go test -v -run TestCreateCampaign
func TestCreateCampaign(t *testing.T) {
	t.Parallel()

	client, err := initClient()
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	data := new(Campaign)
	data.Name = "testCampaign1234"
	data.StartTime = DateTime{time.Now()}
	data.Status = CampaignStatusEnabled
	data.BillingEvent = BillingEventTypeTAPS
	data.BudgetAmount = &Money{
		Amount:   "100.00",
		Currency: "USD",
	}
	data.DailyBudgetAmount = &Money{
		Amount:   "10.00",
		Currency: "USD",
	}
	data.AdamID = 6476530741
	data.CountriesOrRegions = []string{
		"US",
	}
	data.SupplySources = []CampaignSupplySource{
		"APPSTORE_TODAY_TAB",
	}
	data.AdChannelType = CampaignAdChannelTypeDisplay
	res, err := client.Campaigns.CreateCampaign(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(res.Campaign)
}
