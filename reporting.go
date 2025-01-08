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

package asa

import (
	"fmt"
)

// ReportingService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/reports
type ReportingService service

// ReportingRequestGranularity is the report data organized by hour, day, week, and month.
type ReportingRequestGranularity string

const (
	// ReportingRequestGranularityTypeHourly is for a reporting request granularity on Hourly.
	ReportingRequestGranularityTypeHourly ReportingRequestGranularity = "HOURLY"
	// ReportingRequestGranularityTypeDaily is for a reporting request granularity on Daily.
	ReportingRequestGranularityTypeDaily ReportingRequestGranularity = "DAILY"
	// ReportingRequestGranularityTypeWeekly is for a reporting request granularity on Weekly.
	ReportingRequestGranularityTypeWeekly ReportingRequestGranularity = "WEEKLY"
	// ReportingRequestGranularityTypeMonthly is for a reporting request granularity on Monthly.
	ReportingRequestGranularityTypeMonthly ReportingRequestGranularity = "MONTHLY"
)

// ReportingRequestGroupBy is used to group responses by selected dimensions.
type ReportingRequestGroupBy string

const (
	// ReportingRequestGroupByTypeAdminArea is for a reporting request group by on adminArea.
	ReportingRequestGroupByTypeAdminArea ReportingRequestGroupBy = "adminArea"
	// ReportingRequestGroupByTypeAgeRange is for a reporting request group by on ageRange.
	ReportingRequestGroupByTypeAgeRange ReportingRequestGroupBy = "ageRange"
	// ReportingRequestGroupByTypeCountryCode is for a reporting request group by on countryCode.
	ReportingRequestGroupByTypeCountryCode ReportingRequestGroupBy = "countryCode"
	// ReportingRequestGroupByTypeCountryOrRegion is for a reporting request group by on countryOrRegion.
	ReportingRequestGroupByTypeCountryOrRegion ReportingRequestGroupBy = "countryOrRegion"
	// ReportingRequestGroupByTypeDeviceClass is for a reporting request group by on deviceClass.
	ReportingRequestGroupByTypeDeviceClass ReportingRequestGroupBy = "deviceClass"
	// ReportingRequestGroupByTypeGender is for a reporting request group by on gender.
	ReportingRequestGroupByTypeGender ReportingRequestGroupBy = "gender"
	// ReportingRequestGroupByTypeLocality is for a reporting request group by on locality.
	ReportingRequestGroupByTypeLocality ReportingRequestGroupBy = "locality"
)

// ReportingKeywordMatchType is an automated keyword and bidding strategy.
type ReportingKeywordMatchType string

const (
	// ReportingKeywordMatchTypeAuto Use this value to specify that the system serves impressions with optimized keywords, in addition to those you explicitly add to the ad group.
	ReportingKeywordMatchTypeAuto ReportingKeywordMatchType = "AUTO"
	// ReportingKeywordMatchTypeExact Use this value to ensure your ads don’t run on relevant, close variants of a keyword, such as singulars, plurals, misspellings, synonyms, related searches, and phrases that include that term.
	ReportingKeywordMatchTypeExact ReportingKeywordMatchType = "EXACT"
	// ReportingKeywordMatchTypeBroad Use this value for the most control over searches your ad may appear in. You can target a specific term and its close variants, such as common misspellings and plurals. Your ad may receive fewer impressions as a result, but your tap-through rates (TTRs) and conversions on those impressions may be higher because you’re reaching users most interested in your app.
	ReportingKeywordMatchTypeBroad ReportingKeywordMatchType = "BROAD"
)

// SearchTermSource is the source of the keyword to use as a search term.
type SearchTermSource string

const (
	// SearchTermSourceAuto is the value to use to ensure Search Match automatically matches your ads.
	SearchTermSourceAuto SearchTermSource = "AUTO"
	// SearchTermSourceTargeted is a bidded keyword.
	SearchTermSourceTargeted SearchTermSource = "TARGETED"
)

// Row is the report metrics organized by time granularity.
//
// https://developer.apple.com/documentation/apple_search_ads/row
type Row struct {
	Insights    *InsightsObject     `json:"insights,omitempty"`
	Granularity []*ExtendedSpendRow `json:"granularity,omitempty"`
	Metadata    *MetaDataObject     `json:"metadata,omitempty"`
	Other       bool                `json:"other,omitempty"`
	Total       *SpendRow           `json:"total,omitempty"`
}

// CampaignAppDetail is the app data to fetch from campaign-level reports
//
// https://developer.apple.com/documentation/apple_search_ads/campaignappdetail
type CampaignAppDetail struct {
	AppName string `json:"appName"`
	AdamID  int64  `json:"adamId"`
}

// MetaDataObject is the report response objects
//
// https://developer.apple.com/documentation/apple_search_ads/metadataobject
type MetaDataObject struct {
	AdGroupID                          int64                                       `json:"adGroupID,omitempty"`
	AdGroupName                        string                                      `json:"adGroupName,omitempty"`
	AdGroupDeleted                     bool                                        `json:"adGroupDeleted,omitempty"`
	CampaignID                         int64                                       `json:"campaignId,omitempty"`
	CampaignName                       string                                      `json:"campaignName,omitempty"`
	Deleted                            bool                                        `json:"deleted,omitempty"`
	CampaignStatus                     CampaignStatus                              `json:"campaignStatus,omitempty"`
	App                                *CampaignAppDetail                          `json:"app,omitempty"`
	ServingStatus                      CampaignServingStatus                       `json:"servingStatus,omitempty"`
	ServingStateReasons                []CampaignServingStateReason                `json:"servingStateReasons,omitempty"`
	CountriesOrRegions                 []string                                    `json:"countriesOrRegions,omitempty"`
	ModificationTime                   DateTime                                    `json:"modificationTime,omitempty"`
	TotalBudget                        *Money                                      `json:"totalBudget,omitempty"`
	DailyBudget                        *Money                                      `json:"dailyBudget,omitempty"`
	DisplayStatus                      CampaignDisplayStatus                       `json:"displayStatus,omitempty"`
	SupplySources                      []CampaignSupplySource                      `json:"supplySources,omitempty"`
	AdChannelType                      CampaignAdChannelType                       `json:"adChannelType,omitempty"`
	OrgID                              int                                         `json:"orgId,omitempty"`
	CountryOrRegionServingStateReasons *CampaignCountryOrRegionServingStateReasons `json:"countryOrRegionServingStateReasons,omitempty"`
	BillingEvent                       string                                      `json:"billingEvent,omitempty"`
	KeywordID                          int64                                       `json:"keywordID,omitempty"`
	Keyword                            string                                      `json:"keyword,omitempty"`
	KeywordStatus                      string                                      `json:"keywordStatus,omitempty"`
	KeywordDisplayStatus               string                                      `json:"keywordDisplayStatus,omitempty"`
	MatchType                          *ReportingKeywordMatchType                  `json:"matchType,omitempty"`
	CountryOrRegion                    string                                      `json:"countryOrRegion,omitempty"`
	SearchTermText                     *string                                     `json:"SearchTermText,omitempty"`
	SearchTermSource                   *SearchTermSource                           `json:"searchTermSource,omitempty"`
}

// GrandTotalsRow is the summary of cumulative metrics
//
// https://developer.apple.com/documentation/apple_search_ads/grandtotalsrow
type GrandTotalsRow struct {
	Other bool      `json:"other,omitempty"`
	Total *SpendRow `json:"total,omitempty"`
}

// SpendRow is the reporting response metrics
//
// https://developer.apple.com/documentation/apple_search_ads/spendrow
type SpendRow struct {
	AvgCPM            *Money  `json:"avgCPM,omitempty"`
	AvgCPT            *Money  `json:"avgCPT,omitempty"`
	Impressions       int64   `json:"impressions,omitempty"`
	LocalSpend        *Money  `json:"localSpend,omitempty"`
	TapInstallCPI     *Money  `json:"tapInstallCPI,omitempty"`
	TapInstallRate    float64 `json:"tapInstallRate,omitempty"`
	TapInstalls       int64   `json:"tapInstalls,omitempty"`
	TapNewDownloads   int64   `json:"tapNewDownloads,omitempty"`
	TapReDownloads    int64   `json:"tapRedownloads,omitempty"`
	Taps              int64   `json:"taps,omitempty"`
	TotalAvgCPI       *Money  `json:"totalAvgCPI,omitempty"`
	TotalInstallRate  float64 `json:"totalInstallRate,omitempty"`
	TotalInstalls     int64   `json:"totalInstalls,omitempty"`
	TotalNewDownloads int64   `json:"totalNewDownloads,omitempty"`
	TotalReDownloads  int64   `json:"totalRedownloads,omitempty"`
	Ttr               float64 `json:"ttr,omitempty"`
	ViewInstalls      int64   `json:"viewInstalls,omitempty"`
	ViewNewDownloads  int64   `json:"viewNewDownloads,omitempty"`
	ViewReDownloads   int64   `json:"viewRedownloads,omitempty"`
}

// ExtendedSpendRow is the descriptions of metrics with dates
//
// https://developer.apple.com/documentation/apple_search_ads/extendedspendrow
type ExtendedSpendRow struct {
	AvgCPM            *Money  `json:"avgCPM,omitempty"`
	AvgCPT            *Money  `json:"avgCPT,omitempty"`
	Date              Date    `json:"date,omitempty"`
	Impressions       int64   `json:"impressions,omitempty"`
	LocalSpend        *Money  `json:"localSpend,omitempty"`
	TapInstallCPI     *Money  `json:"tapInstallCPI,omitempty"`
	TapInstallRate    float64 `json:"tapInstallRate,omitempty"`
	TapInstalls       int64   `json:"tapInstalls,omitempty"`
	TapNewDownloads   int64   `json:"tapNewDownloads,omitempty"`
	TapReDownloads    int64   `json:"tapRedownloads,omitempty"`
	Taps              int64   `json:"taps,omitempty"`
	TotalAvgCPI       *Money  `json:"totalAvgCPI,omitempty"`
	TotalInstallRate  float64 `json:"totalInstallRate,omitempty"`
	TotalInstalls     int64   `json:"totalInstalls,omitempty"`
	TotalNewDownloads int64   `json:"totalNewDownloads,omitempty"`
	TotalReDownloads  int64   `json:"totalRedownloads,omitempty"`
	Ttr               float64 `json:"ttr,omitempty"`
	ViewInstalls      int64   `json:"viewInstalls,omitempty"`
	ViewNewDownloads  int64   `json:"viewNewDownloads,omitempty"`
	ViewReDownloads   int64   `json:"viewRedownloads,omitempty"`
}

// InsightsObject is a parent object for bid recommendations
//
// https://developer.apple.com/documentation/apple_search_ads/insightsobject
type InsightsObject struct {
	BidRecommendation *KeywordBidRecommendation `json:"bidRecommendation,omitempty"`
}

// KeywordBidRecommendation is the bid recommendation range for a keyword
//
// https://developer.apple.com/documentation/apple_search_ads/keywordbidrecommendation
type KeywordBidRecommendation struct {
	SuggestedBidAmount *Money `json:"suggestedBidAmount,omitempty"`
}

// ReportingRequest is the report request body
//
// https://developer.apple.com/documentation/apple_search_ads/reportingrequest
type ReportingRequest struct {
	StartTime                  ReqDate                     `json:"startTime,omitempty"`
	EndTime                    ReqDate                     `json:"endTime,omitempty"`
	Granularity                ReportingRequestGranularity `json:"granularity,omitempty"`
	TimeZone                   ReportingRequestTimeZone    `json:"timeZone,omitempty"`
	GroupBy                    []ReportingRequestGroupBy   `json:"groupBy,omitempty"`
	ReturnGrandTotals          bool                        `json:"returnGrandTotals,omitempty"`
	ReturnRecordsWithNoMetrics bool                        `json:"returnRecordsWithNoMetrics,omitempty"`
	ReturnRowTotals            bool                        `json:"returnRowTotals,omitempty"`
	Selector                   *Selector                   `json:"selector,omitempty"`
}

// ReportingResponseBody is a container for the report response body
//
// https://developer.apple.com/documentation/apple_search_ads/reportingresponsebody
type ReportingResponseBody struct {
	ReportingCampaign *ReportingResponse `json:"data,omitempty"`
	Pagination        *PageDetail        `json:"pagination,omitempty"`
	Error             *ErrorResponseBody `json:"error,omitempty"`
}

// ReportingResponse is a container for report metrics
//
// https://developer.apple.com/documentation/apple_search_ads/reportingresponse
type ReportingResponse struct {
	ReportingDataResponse *ReportingDataResponse `json:"reportingDataResponse,omitempty"`
}

// ReportingDataResponse is the total metrics for a report
//
// https://developer.apple.com/documentation/apple_search_ads/reportingdataresponse
type ReportingDataResponse struct {
	GrandTotals *GrandTotalsRow `json:"grandTotals,omitempty"`
	Rows        []Row           `json:"row,omitempty"`
}

// GetCampaignLevelReports fetches reports for campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/get_campaign-level_reports
func (s *ReportingService) GetCampaignLevelReports(params *ReportingRequest) (*ReportingResponseBody, error) {
	url := "reports/campaigns"
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}

// GetAdGroupLevelReports fetches reports for ad groups within a campaig
//
// https://developer.apple.com/documentation/apple_search_ads/get_ad_group-level_reports
func (s *ReportingService) GetAdGroupLevelReports(campaignID int64, params *ReportingRequest) (*ReportingResponseBody, error) {
	url := fmt.Sprintf("reports/campaigns/%d/adgroups", campaignID)
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}

// GetKeywordLevelReports fetches reports for targeting keywords within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_keyword-level_reports
func (s *ReportingService) GetKeywordLevelReports(campaignID int64, params *ReportingRequest) (*ReportingResponseBody, error) {
	url := fmt.Sprintf("reports/campaigns/%d/keywords", campaignID)
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}

// GetAdGroupKeywordLevelReports Fetches reports for targeting keywords within an ad group.
//
// https://developer.apple.com/documentation/apple_search_ads/get_keyword-level_within_ad_group_reports
func (s *ReportingService) GetAdGroupKeywordLevelReports(campaignID, adGroupID int64, params *ReportingRequest) (*ReportingResponseBody, error) {
	url := fmt.Sprintf("reports/campaigns/%d/adgroups/%d/keywords", campaignID, adGroupID)
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}

// GetSearchTermLevelReports fetches reports for search terms within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_search_term-level_reports
func (s *ReportingService) GetSearchTermLevelReports(campaignID int64, params *ReportingRequest) (*ReportingResponseBody, error) {
	url := fmt.Sprintf("reports/campaigns/%d/searchterms", campaignID)
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}

// GetAdGroupSearchTermLevelReports Fetches reports for search terms within an ad group.
//
// https://developer.apple.com/documentation/apple_search_ads/get_search_term-level_within_ad_group_reports
func (s *ReportingService) GetAdGroupSearchTermLevelReports(campaignID, adGroupID int64, params *ReportingRequest) (*ReportingResponseBody, error) {
	url := fmt.Sprintf("reports/campaigns/%d/adgroups/%d/searchterms", campaignID, adGroupID)
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}

// GetCreativeSetLevelReports fetches reports for Creative Sets within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_creative_set-level_reports
func (s *ReportingService) GetCreativeSetLevelReports(campaignID int64, params *ReportingRequest) (*ReportingResponseBody, error) {
	url := fmt.Sprintf("reports/campaigns/%d/creativesets", campaignID)
	res := new(ReportingResponseBody)
	err := s.client.post(url, res, params)

	return res, err
}
