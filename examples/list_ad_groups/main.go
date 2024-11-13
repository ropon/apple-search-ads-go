/**
Copyright (C) 2021 Mehmet Gungoren.
This file is part of apple-search-ads-go, a package for working with Apple's
Search Ads API.
apple-search-ads-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
apple-search-ads-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with apple-search-ads-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ropon/apple-search-ads-go"
	"github.com/ropon/apple-search-ads-go/examples/util"
	"log"
)

var (
	campaignName = flag.String("campaign_name", "", "Campaign Name for an Campaign")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the Apple Search Ads client
	client := asa.NewClient(auth.Client())

	campaign, err := util.GetCampaign(ctx, client, &asa.Selector{
		Conditions: []*asa.Condition{
			{
				Field:    "Name",
				Operator: asa.ConditionOperatorEquals,
				Values:   []string{*campaignName},
			},
		},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	params := &asa.GetAllAdGroupsQuery{}
	for {
		adGroupsResponse, _, err := client.AdGroups.GetAllAdGroups(ctx, campaign.ID, params)
		if err != nil {
			log.Fatal(err)
		}

		for _, adGroup := range adGroupsResponse.AdGroups {
			fmt.Println(adGroup.Name)
		}
		pageDetail := adGroupsResponse.Pagination
		lastOffset := pageDetail.StartIndex + len(adGroupsResponse.AdGroups)
		if lastOffset < pageDetail.TotalResults {
			params.Offset += int32(len(adGroupsResponse.AdGroups))
		} else {
			break
		}
	}
}
