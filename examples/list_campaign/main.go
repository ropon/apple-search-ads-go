package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ropon/apple-search-ads-go/asa"
	"github.com/ropon/apple-search-ads-go/examples/util"
	"log"
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
	req := &asa.GetAllCampaignQuery{
		Offset: 0,
		Limit:  20,
	}
	res, _, err := client.Campaigns.GetAllCampaigns(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(res)
}
