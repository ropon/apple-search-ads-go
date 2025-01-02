package asa

import "fmt"

// AppService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/search_apps_and_geolocations
type AppService service

// AppInfo is the response to an app search request
//
// https://developer.apple.com/documentation/apple_search_ads/appinfo
type AppInfo struct {
	AdamID               int64    `json:"adamId,omitempty"`
	AppName              string   `json:"appName,omitempty"`
	CountryOrRegionCodes []string `json:"countryOrRegionCodes,omitempty"`
	DeveloperName        string   `json:"developerName"`
}

// SearchAppsQuery defines query parameter for SearchApps endpoint.
type SearchAppsQuery struct {
	Limit           int32  `form:"limit,omitempty"`
	Offset          int32  `form:"offset,omitempty"`
	Query           string `form:"query,omitempty"`
	ReturnOwnedApps bool   `form:"returnOwnedApps,omitempty"`
}

// AppInfoListResponse is the response details of app search requests
//
// https://developer.apple.com/documentation/apple_search_ads/appinfolistresponse
type AppInfoListResponse struct {
	AppInfos   []*AppInfo         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// EligibilityRecord App eligibility parameters that an API response returns.
//
// https://developer.apple.com/documentation/apple_search_ads/eligibilityrecord
type EligibilityRecord struct {
	AdamID          int64  `json:"adamId,omitempty"`
	CountryOrRegion string `json:"countryOrRegion,omitempty"`
	DeviceClass     string `json:"deviceClass,omitempty"`
	MinAge          int32  `json:"minAge,omitempty"`
	State           string `json:"state,omitempty"`
	SupplySource    string `json:"supplySource,omitempty"`
}

// EligibilityRecordListResponse The response details to an app eligibility request.
//
// https://developer.apple.com/documentation/apple_search_ads/eligibilityrecordlistresponse
type EligibilityRecordListResponse struct {
	EligibilityRecords []*EligibilityRecord `json:"data,omitempty"`
	Pagination         *PageDetail          `json:"pagination,omitempty"`
}

// SearchApps Searches for iOS apps to promote in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/search_for_ios_apps
func (s *AppService) SearchApps(params *SearchAppsQuery) (*AppInfoListResponse, error) {
	url := "search/apps"
	res := new(AppInfoListResponse)
	err := s.client.get(url, res, params)

	return res, err
}

// FindAppEligibilityRecords Fetches app eligibility records by adam ID.
//
// https://developer.apple.com/documentation/apple_search_ads/find_app_eligibility_records
func (s *AppService) FindAppEligibilityRecords(adamId int64, selector *Selector) (*EligibilityRecordListResponse, error) {
	url := fmt.Sprintf("apps/%d/eligibilities/find", adamId)
	res := new(EligibilityRecordListResponse)
	err := s.client.post(url, res, selector)

	return res, err
}
