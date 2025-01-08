package asa

import (
	"fmt"
)

// AdGroupService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/ad_groups
type AdGroupService service

// AdGroupDisplayStatus defines model for AdGroupDisplayStatus.
//
// https://developer.apple.com/documentation/apple_search_ads/adgroup
type AdGroupDisplayStatus string

const (
	// AdGroupDisplayStatusDelete is for an ad group display status on Deleted.
	AdGroupDisplayStatusDelete AdGroupDisplayStatus = "DELETED"
	// AdGroupDisplayStatusOnHold is for an ad group display status on On Hold.
	AdGroupDisplayStatusOnHold AdGroupDisplayStatus = "ON_HOLD"
	// AdGroupDisplayStatusPaused is for an ad group display status on Paused.
	AdGroupDisplayStatusPaused AdGroupDisplayStatus = "PAUSED"
	// AdGroupDisplayStatusRunning is for an ad group display status on Running.
	AdGroupDisplayStatusRunning AdGroupDisplayStatus = "RUNNING"
)

// AdGroupPaymentModel defines model for AdGroupPaymentModel.
//
// https://developer.apple.com/documentation/apple_search_ads/paymentmodel
type AdGroupPaymentModel string

const (
	// AdGroupPaymentModelPAYG is for an ad group payment model PAYG.
	AdGroupPaymentModelPAYG AdGroupPricingModel = "PAYG"
	// AdGroupPaymentModelLOC is for an ad group payment model LOC.
	AdGroupPaymentModelLOC AdGroupPricingModel = "LOC"
)

// AdGroupPricingModel defines model for AdGroupPricingModel.
//
// https://developer.apple.com/documentation/apple_search_ads/pricingmodel
type AdGroupPricingModel string

const (
	// AdGroupPricingModelCPC is for an ad group pricing model CPC.
	AdGroupPricingModelCPC AdGroupPricingModel = "CPC"
	// AdGroupPricingModelCPM is for an ad group pricing model CPM.
	AdGroupPricingModelCPM AdGroupPricingModel = "CPM"
)

// ServingStateReason is that displays when an ad group isn’t running.
type ServingStateReason string

const (
	// ServingStateReasonAdGroupPausedByUser is for an ad group serving state reason Ad Group Paused By User.
	ServingStateReasonAdGroupPausedByUser ServingStateReason = "AD_GROUP_PAUSED_BY_USER"
	// ServingStateReasonAdGroupEndDateReached is for an ad group serving state reason Ad Group End Date Reached.
	ServingStateReasonAdGroupEndDateReached ServingStateReason = "ADGROUP_END_DATE_REACHED"
	// ServingStateReasonAppNotSupport is for an ad group serving state reason App Not Support.
	ServingStateReasonAppNotSupport ServingStateReason = "APP_NOT_SUPPORT"
	// ServingStateReasonAudienceBelowThreshold is for an ad group serving state reason Audience Below Threshold.
	ServingStateReasonAudienceBelowThreshold ServingStateReason = "AUDIENCE_BELOW_THRESHOLD"
	// ServingStateReasonCampaignNotRunning is for an ad group serving state reason Campaign Not Running.
	ServingStateReasonCampaignNotRunning ServingStateReason = "CAMPAIGN_NOT_RUNNING"
	// ServingStateReasonDeletedByUser is for an ad group serving state reason Deleted By User.
	ServingStateReasonDeletedByUser ServingStateReason = "DELETED_BY_USER"
	// ServingStateReasonPendingAudienceVerification is for an ad group serving state reason Pending Audience Verification.
	ServingStateReasonPendingAudienceVerification ServingStateReason = "PENDING_AUDIENCE_VERIFICATION"
	// ServingStateReasonStartDateInTheFuture is for an ad group serving state reason Start Date in The Future.
	ServingStateReasonStartDateInTheFuture ServingStateReason = "START_DATE_IN_THE_FUTURE"
)

// AdGroupServingStatus is the status of whether the ad group is serving.
type AdGroupServingStatus string

const (
	// AdGroupServingStatusNotRunning is for an ad group serving status Not Running.
	AdGroupServingStatusNotRunning AdGroupServingStatus = "NOT_RUNNING"
	// AdGroupServingStatusRunning is for an ad group serving status Running.
	AdGroupServingStatusRunning AdGroupServingStatus = "RUNNING"
)

// AdGroupStatus is the user-controlled status to enable or pause the ad group.
type AdGroupStatus string

const (
	// AdGroupStatusEnabled is for an ad group status Enabled.
	AdGroupStatusEnabled AdGroupStatus = "ENABLED"
	// AdGroupStatusPaused is for an ad group status Paused.
	AdGroupStatusPaused AdGroupStatus = "PAUSED"
)

// AdGroupDeviceClass is targeting criteria values for device class targeting.
type AdGroupDeviceClass string

const (
	// AdGroupDeviceClassIpad is for ad group targeting criteria values for Ipad.
	AdGroupDeviceClassIpad AdGroupDeviceClass = "IPAD"
	// AdGroupDeviceClassIphone is for ad group targeting criteria values for Iphone.
	AdGroupDeviceClassIphone AdGroupDeviceClass = "IPHONE"
)

// AdGroupGender is the targeting criteria values for gender.
type AdGroupGender string

const (
	// AdGroupGenderFemale is the targeting gender criteria for Female.
	AdGroupGenderFemale AdGroupGender = "F"
	// AdGroupGenderMale is the targeting gender criteria for Male.
	AdGroupGenderMale AdGroupGender = "M"
)

// ConditionOperator is the operator values compare attributes to a list of specified values.
type ConditionOperator string

const (
	// ConditionOperatorBetween is the attribute matches the values within a specified range. The values can be numbers, text, or dates.
	ConditionOperatorBetween ConditionOperator = "BETWEEN"
	// ConditionOperatorContains is the attribute matches the value in the specified list.
	ConditionOperatorContains ConditionOperator = "CONTAINS"
	// ConditionOperatorContainsAll is the attribute has all of the values in the specified list.
	ConditionOperatorContainsAll ConditionOperator = "CONTAINS_ALL"
	// ConditionOperatorContainsAny is the attribute contains any of the values in the specified list.
	ConditionOperatorContainsAny ConditionOperator = "CONTAINS_ANY"
	// ConditionOperatorEndsWith is the attribute matches the suffix of a string.
	ConditionOperatorEndsWith ConditionOperator = "ENDSWITH"
	// ConditionOperatorEquals is the attribute contains exact values.
	ConditionOperatorEquals ConditionOperator = "EQUALS"
	// ConditionOperatorGreaterThan is the value is greater than the specified value.
	ConditionOperatorGreaterThan ConditionOperator = "GREATER_THAN"
	// ConditionOperatorLessThan is the value is less than the specified value.
	ConditionOperatorLessThan ConditionOperator = "LESS_THAN"
	// ConditionOperatorStartsWith is the attribute matches the prefix of a string.
	ConditionOperatorStartsWith ConditionOperator = "STARTSWITH"
	// ConditionOperatorIn is the attribute matches any value in a list of specified values.
	ConditionOperatorIn ConditionOperator = "IN"
	// ConditionOperatorLike is the attribute like the value in the specified value.
	ConditionOperatorLike ConditionOperator = "LIKE"
	// ConditionOperatorNotEqual is the attribute not contains exact values.
	ConditionOperatorNotEqual ConditionOperator = "NOT_EQUALS"
	// ConditionOperatorIs is the attribute contains any of the values in the specified list.
	ConditionOperatorIs ConditionOperator = "IS"
)

// SortOrder is the order of grouped results.
type SortOrder string

const (
	// SortingOrderAscending is for sort order of Ascending.
	SortingOrderAscending SortOrder = "ASCENDING"
	// SortingOrderDescending is for sort order of Descending.
	SortingOrderDescending SortOrder = "DESCENDING"
)

// AdGroup is the response to ad group requests
//
// https://developer.apple.com/documentation/apple_search_ads/adgroup
type AdGroup struct {
	AutomatedKeywordsOptIn bool                 `json:"automatedKeywordsOptIn,omitempty"`
	CampaignID             int64                `json:"campaignID,omitempty"`
	CpaGoal                *Money               `json:"cpaGoal,omitempty"`
	CreationTime           DateTime             `json:"creationTime,omitempty"`
	DefaultBidAmount       *Money               `json:"defaultBidAmount,omitempty"`
	Deleted                bool                 `json:"deleted,omitempty"`
	DisplayStatus          AdGroupDisplayStatus `json:"displayStatus,omitempty"`
	EndTime                *DateTime            `json:"endTime,omitempty"`
	ID                     int64                `json:"id,omitempty"`
	ModificationTime       DateTime             `json:"modificationTime,omitempty"`
	Name                   string               `json:"name,omitempty"`
	OrgID                  int64                `json:"orgId,omitempty"`
	PaymentModel           AdGroupPaymentModel  `json:"paymentModel,omitempty"`
	PricingModel           AdGroupPricingModel  `json:"pricingModel,omitempty"`
	ServingStateReasons    []ServingStateReason `json:"servingStateReasons,omitempty"`
	ServingStatus          AdGroupServingStatus `json:"servingStatus,omitempty"`
	StartTime              DateTime             `json:"startTime,omitempty"`
	Status                 AdGroupStatus        `json:"status,omitempty"`
	TargetingDimensions    *TargetingDimensions `json:"targetingDimensions,omitempty"`
}

// Money is the response to requests for budget amounts in campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/money
type Money struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// TargetingDimensions is the criteria to use with ad groups to narrow the audience that views the ads
//
// https://developer.apple.com/documentation/apple_search_ads/targetingdimensions
type TargetingDimensions struct {
	AdminArea      *AdminAreaCriteria     `json:"adminArea,omitempty"`
	Age            *AgeCriteria           `json:"age,omitempty"`
	AppCategories  *AppCategoryCriteria   `json:"appCategories,omitempty"`
	AppDownloaders *AppDownloaderCriteria `json:"appDownloaders,omitempty"`
	Country        *CountryCriteria       `json:"country,omitempty"`
	DayPart        *DayPartCriteria       `json:"daypart,omitempty"`
	DeviceClass    *DeviceClassCriteria   `json:"deviceClass,omitempty"`
	Gender         *GenderCriteria        `json:"gender,omitempty"`
	Locality       *LocalityCriteria      `json:"locality,omitempty"`
}

// AdminAreaCriteria is the defined targeted audience by administrative area
//
// https://developer.apple.com/documentation/apple_search_ads/adminareacriteria
type AdminAreaCriteria struct {
	Included []string `json:"included,omitempty"`
}

// AgeCriteria is the defined targeted audience to include using the age demographic
//
// https://developer.apple.com/documentation/apple_search_ads/agecriteria
type AgeCriteria struct {
	Included []*AgeRange `json:"included,omitempty"`
}

// AgeRange is the defined target audience to include using the age range demographic
//
// https://developer.apple.com/documentation/apple_search_ads/agerange
type AgeRange struct {
	MaxAge int32 `json:"maxAge,omitempty"`
	MinAge int32 `json:"minAge,omitempty"`
}

// AppCategoryCriteria The defined target audience by app category.
//
// https://developer.apple.com/documentation/apple_search_ads/appcategorycriteria
type AppCategoryCriteria struct {
	Excluded []int64 `json:"excluded,omitempty"`
	Included []int64 `json:"included,omitempty"`
}

// AppDownloaderCriteria is the defined targeted audience according to app downloads
//
// https://developer.apple.com/documentation/apple_search_ads/appdownloadercriteria
type AppDownloaderCriteria struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

// CountryCriteria is the defined targeted audience by country or region
//
// https://developer.apple.com/documentation/apple_search_ads/countrycriteria
type CountryCriteria struct {
	Included []string `json:"included,omitempty"`
}

// DayPartCriteria is the defined targeted audience to include for a specific time of day
//
// https://developer.apple.com/documentation/apple_search_ads/daypartcriteria
type DayPartCriteria struct {
	UserTime *DaypartDetail `json:"userTime,omitempty"`
}

// DaypartDetail is the defined targeted audience to include by a specific time of day
//
// https://developer.apple.com/documentation/apple_search_ads/daypartdetail
type DaypartDetail struct {
	Included []int32 `json:"included,omitempty"`
}

// DeviceClassCriteria is the defined targeted audience to include by device type
//
// https://developer.apple.com/documentation/apple_search_ads/deviceclasscriteria
type DeviceClassCriteria struct {
	Included []AdGroupDeviceClass `json:"included,omitempty"`
}

// GenderCriteria is the defined targeted audience to include using the gender demographic
//
// https://developer.apple.com/documentation/apple_search_ads/gendercriteria
type GenderCriteria struct {
	Included []AdGroupGender `json:"included,omitempty"`
}

// LocalityCriteria is the defined targeted audience by locality
//
// https://developer.apple.com/documentation/apple_search_ads/localitycriteria
type LocalityCriteria struct {
	Included []string `json:"included,omitempty"`
}

// GetAllAdGroupsQuery defines query parameter for GetAllAdGroups endpoint.
type GetAllAdGroupsQuery struct {
	Limit  int32 `form:"limit,omitempty"`
	Offset int32 `form:"offset,omitempty"`
}

// Selector is the selector objects available to filter returned data
//
// https://developer.apple.com/documentation/apple_search_ads/selector
type Selector struct {
	Conditions []*Condition `json:"conditions,omitempty"`
	Fields     []string     `json:"fields,omitempty"`
	OrderBy    []*Sorting   `json:"orderBy,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

// Condition is the list of condition objects that allow users to filter a list of records
//
// https://developer.apple.com/documentation/apple_search_ads/condition
type Condition struct {
	Field    string            `json:"field,omitempty"`
	Operator ConditionOperator `json:"operator,omitempty"`
	Values   []string          `json:"values,omitempty"`
}

// Sorting is the order of grouped results
//
// https://developer.apple.com/documentation/apple_search_ads/sorting
type Sorting struct {
	Field     string    `json:"field,omitempty"`
	SortOrder SortOrder `json:"sortOrder,omitempty"`
}

// Pagination is the procedure to refine returned results using limit and offset parameters
//
// https://developer.apple.com/documentation/apple_search_ads/pagination
type Pagination struct {
	Limit  uint32 `json:"limit,omitempty"`
	Offset uint32 `json:"offset,omitempty"`
}

// AdGroupUpdateRequest https://developer.apple.com/documentation/apple_search_ads/adgroupupdate
type AdGroupUpdateRequest struct {
	AutomatedKeywordsOptIn bool                 `json:"automatedKeywordsOptIn,omitempty"`
	CpaGoal                *Money               `json:"cpaGoal,omitempty"`
	DefaultBidAmount       *Money               `json:"defaultBidAmount,omitempty"`
	EndTime                DateTime             `json:"endTime,omitempty"`
	Name                   string               `json:"name,omitempty"`
	StartTime              DateTime             `json:"startTime,omitempty"`
	Status                 AdGroupStatus        `json:"status,omitempty"`
	TargetingDimensions    *TargetingDimensions `json:"targetingDimensions"`
}

// AdGroupResponse is a container for the ad group response body
//
// https://developer.apple.com/documentation/apple_search_ads/adgroupresponse
type AdGroupResponse struct {
	AdGroup    *AdGroup           `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// AdGroupListResponse is the response details of ad group requests
//
// https://developer.apple.com/documentation/apple_search_ads/adgrouplistresponse
type AdGroupListResponse struct {
	AdGroups   []*AdGroup         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// FindAdGroups fetches ad groups within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/find_ad_groups
func (s *AdGroupService) FindAdGroups(campaignID int64, selector *Selector) (*AdGroupListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/find", campaignID)
	res := new(AdGroupListResponse)
	err := s.client.post(url, res, selector)
	return res, err
}

// GetAdGroup fetches a specific ad group with a campaign and ad group identifier
//
// https://developer.apple.com/documentation/apple_search_ads/get_an_ad_group
func (s *AdGroupService) GetAdGroup(campaignID int64, adGroupID int64) (*AdGroupResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, adGroupID)
	res := new(AdGroupResponse)
	err := s.client.get(url, res)
	return res, err
}

// GetAllAdGroups fetches all ad groups with a campaign identifier.
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_ad_groups
func (s *AdGroupService) GetAllAdGroups(campaignID int64, params *GetAllAdGroupsQuery) (*AdGroupListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups", campaignID)
	res := new(AdGroupListResponse)
	err := s.client.get(url, res, params)

	return res, err
}

// CreateAdGroup creates an ad group as part of a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/create_an_ad_group
func (s *AdGroupService) CreateAdGroup(campaignID int64, adGroup *AdGroup) (*AdGroupResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups", campaignID)
	res := new(AdGroupResponse)
	err := s.client.post(url, res, adGroup)

	return res, err
}

// DeleteAdGroup deletes an ad group with a campaign and ad group identifier.
//
// https://developer.apple.com/documentation/apple_search_ads/delete_an_adgroup
func (s *AdGroupService) DeleteAdGroup(campaignID int64, adGroupID int64) (*BaseResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, adGroupID)
	res := new(BaseResponse)
	err := s.client.delete(url, res)

	return res, err
}

// UpdateAdGroup updates an ad group with an ad group identifier.
//
// https://developer.apple.com/documentation/apple_search_ads/update_an_ad_group
func (s *AdGroupService) UpdateAdGroup(campaignID int64, adGroupID int64, req *AdGroupUpdateRequest) (*AdGroupResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, adGroupID)
	res := new(AdGroupResponse)
	err := s.client.put(url, res, req)

	return res, err
}
