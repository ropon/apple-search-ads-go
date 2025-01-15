package asa

import (
	"fmt"
)

// CampaignService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/campaigns
type CampaignService service

// CampaignAdChannelType is the channel type of ad in a campaign.
type CampaignAdChannelType string

const (
	// CampaignAdChannelTypeSearch When supplySources is APPSTORE_SEARCH_RESULTS, the adChannelType must be SEARCH.
	CampaignAdChannelTypeSearch CampaignAdChannelType = "SEARCH"
	// CampaignAdChannelTypeDisplay When supplySources is APPSTORE_SEARCH_TAB, the adChannelType must be DISPLAY.
	CampaignAdChannelTypeDisplay CampaignAdChannelType = "DISPLAY"
)

// BillingEventType is the billing event for a campaign.
type BillingEventType string

const (
	// BillingEventTypeTAPS When the supplySources value is APPSTORE_SEARCH_RESULTS or APPSTORE_SEARCH_TAB, the billingEvent must be TAPS.
	BillingEventTypeTAPS BillingEventType = "TAPS"
	// BillingEventTypeIMPRESSIONS The cost to the advertiser is per impression served.
	BillingEventTypeIMPRESSIONS BillingEventType = "IMPRESSIONS"
)

// CampaignDisplayStatus is the status of the campaign.
type CampaignDisplayStatus string

const (
	// CampaignDisplayStatusRunning is for a campaign status on RUNNING.
	CampaignDisplayStatusRunning CampaignDisplayStatus = "RUNNING"
	// CampaignDisplayStatusOnHold is for a campaign status on ON_HOLD.
	CampaignDisplayStatusOnHold CampaignDisplayStatus = "ON_HOLD"
	// CampaignDisplayStatusPaused is for a campaign status on PAUSED.
	CampaignDisplayStatusPaused CampaignDisplayStatus = "PAUSED"
	// CampaignDisplayStatusDeleted is for a campaign status on DELETED.
	CampaignDisplayStatusDeleted CampaignDisplayStatus = "DELETED"
)

// CampaignServingStateReason is a reason that displays when a campaign can’t run.
type CampaignServingStateReason string

const (
	// CampaignServingStateReasonNoPaymentMethodOnFile is for a campaign serving state reason for NO_PAYMENT_METHOD_ON_FILE.
	CampaignServingStateReasonNoPaymentMethodOnFile CampaignServingStateReason = "NO_PAYMENT_METHOD_ON_FILE"
	// CampaignServingStateReasonMissingBoOrInvoicingFields is for a campaign serving state reason for MISSING_BO_OR_INVOICING_FIELDS.
	CampaignServingStateReasonMissingBoOrInvoicingFields CampaignServingStateReason = "MISSING_BO_OR_INVOICING_FIELDS"
	// CampaignServingStateReasonPausedByUser is for a campaign serving state reason for PAUSED_BY_USER.
	CampaignServingStateReasonPausedByUser CampaignServingStateReason = "PAUSED_BY_USER"
	// CampaignServingStateReasonDeletedByUser is for a campaign serving state reason for DELETED_BY_USER.
	CampaignServingStateReasonDeletedByUser CampaignServingStateReason = "DELETED_BY_USER"
	// CampaignServingStateReasonCampaignEndDateReached is for a campaign serving state reason for CAMPAIGN_END_DATE_REACHED.
	CampaignServingStateReasonCampaignEndDateReached CampaignServingStateReason = "CAMPAIGN_END_DATE_REACHED"
	// CampaignServingStateReasonCampaignStartDateInFuture is for a campaign serving state reason for CAMPAIGN_START_DATE_IN_FUTURE.
	CampaignServingStateReasonCampaignStartDateInFuture CampaignServingStateReason = "CAMPAIGN_START_DATE_IN_FUTURE"
	// CampaignServingStateReasonDailyCapExhausted is for a campaign serving state reason for DAILY_CAP_EXHAUSTED.
	CampaignServingStateReasonDailyCapExhausted CampaignServingStateReason = "DAILY_CAP_EXHAUSTED"
	// CampaignServingStateReasonTotalBudgetExhausted is for a campaign serving state reason for TOTAL_BUDGET_EXHAUSTED.
	CampaignServingStateReasonTotalBudgetExhausted CampaignServingStateReason = "TOTAL_BUDGET_EXHAUSTED"
	// CampaignServingStateReasonCreditCardDeclined is for a campaign serving state reason for CREDIT_CARD_DECLINED.
	CampaignServingStateReasonCreditCardDeclined CampaignServingStateReason = "CREDIT_CARD_DECLINED"
	// CampaignServingStateReasonAppNotEligible is for a campaign serving state reason for APP_NOT_ELIGIBLE.
	CampaignServingStateReasonAppNotEligible CampaignServingStateReason = "APP_NOT_ELIGIBLE"
	// CampaignServingStateReasonAppNotEligibleSearchads is for a campaign serving state reason for APP_NOT_ELIGIBLE_SEARCHADS.
	CampaignServingStateReasonAppNotEligibleSearchads CampaignServingStateReason = "APP_NOT_ELIGIBLE_SEARCHADS"
	// CampaignServingStateReasonAppNotPublishedYet is for a campaign serving state reason for APP_NOT_PUBLISHED_YET.
	CampaignServingStateReasonAppNotPublishedYet CampaignServingStateReason = "APP_NOT_PUBLISHED_YET"
	// CampaignServingStateReasonBoStartDateInFuture is for a campaign serving state reason for BO_START_DATE_IN_FUTURE.
	CampaignServingStateReasonBoStartDateInFuture CampaignServingStateReason = "BO_START_DATE_IN_FUTURE"
	// CampaignServingStateReasonBoEndDateReached is for a campaign serving state reason for BO_END_DATE_REACHED.
	CampaignServingStateReasonBoEndDateReached CampaignServingStateReason = "BO_END_DATE_REACHED"
	// CampaignServingStateReasonBoExhausted is for a campaign serving state reason for BO_EXHAUSTED.
	CampaignServingStateReasonBoExhausted CampaignServingStateReason = "BO_EXHAUSTED"
	// CampaignServingStateReasonOrgPaymentTypeChanged is for a campaign serving state reason for ORG_PAYMENT_TYPE_CHANGED.
	CampaignServingStateReasonOrgPaymentTypeChanged CampaignServingStateReason = "ORG_PAYMENT_TYPE_CHANGED"
	// CampaignServingStateReasonOrgSuspendedPolicyViolation is for a campaign serving state reason for ORG_SUSPENDED_POLICY_VIOLATION.
	CampaignServingStateReasonOrgSuspendedPolicyViolation CampaignServingStateReason = "ORG_SUSPENDED_POLICY_VIOLATION"
	// CampaignServingStateReasonOrgSuspendedFraud is for a campaign serving state reason for ORG_SUSPENDED_FRAUD.
	CampaignServingStateReasonOrgSuspendedFraud CampaignServingStateReason = "ORG_SUSPENDED_FRAUD"
	// CampaignServingStateReasonOrgChargeBackDisputed is for a campaign serving state reason for ORG_CHARGE_BACK_DISPUTED.
	CampaignServingStateReasonOrgChargeBackDisputed CampaignServingStateReason = "ORG_CHARGE_BACK_DISPUTED"
	// CampaignServingStateReasonPausedBySystem is for a campaign serving state reason for PAUSED_BY_SYSTEM.
	CampaignServingStateReasonPausedBySystem CampaignServingStateReason = "PAUSED_BY_SYSTEM"
	// CampaignServingStateReasonLocExhausted is for a campaign serving state reason for LOC_EXHAUSTED.
	CampaignServingStateReasonLocExhausted CampaignServingStateReason = "LOC_EXHAUSTED"
	// CampaignServingStateReasonTaxVerificationPending is for a campaign serving state reason for TAX_VERIFICATION_PENDING.
	CampaignServingStateReasonTaxVerificationPending CampaignServingStateReason = "TAX_VERIFICATION_PENDING"
	// CampaignServingStateReasonSapinLawAgentUnknown is for a campaign serving state reason for SAPIN_LAW_AGENT_UNKNOWN.
	CampaignServingStateReasonSapinLawAgentUnknown CampaignServingStateReason = "SAPIN_LAW_AGENT_UNKNOWN"
	// CampaignServingStateReasonSapinLawFrenchBizUnknown is for a campaign serving state reason for SAPIN_LAW_FRENCH_BIZ_UNKNOWN.
	CampaignServingStateReasonSapinLawFrenchBizUnknown CampaignServingStateReason = "SAPIN_LAW_FRENCH_BIZ_UNKNOWN"
	// CampaignServingStateReasonSapinLawFrenchBiz is for a campaign serving state reason for SAPIN_LAW_FRENCH_BIZ.
	CampaignServingStateReasonSapinLawFrenchBiz CampaignServingStateReason = "SAPIN_LAW_FRENCH_BIZ"
	// CampaignServingStateReasonNoEligibleCountries is for a campaign serving state reason for NO_ELIGIBLE_COUNTRIES.
	CampaignServingStateReasonNoEligibleCountries CampaignServingStateReason = "NO_ELIGIBLE_COUNTRIES"
	// CampaignServingStateReasonAdGroupMissing is for a campaign serving state reason for AD_GROUP_MISSING.
	CampaignServingStateReasonAdGroupMissing CampaignServingStateReason = "AD_GROUP_MISSING"
)

// CampaignSupplySource is the supply source of ads to use in a campaign.
type CampaignSupplySource string

const (
	// CampaignSupplySourceAppstoreSearchResults is for a campaign supply source on APPSTORE_SEARCH_RESULTS.
	CampaignSupplySourceAppstoreSearchResults CampaignSupplySource = "APPSTORE_SEARCH_RESULTS"
	// CampaignSupplySourceAppstoreSearchTab is for a campaign supply source on APPSTORE_SEARCH_TAB.
	CampaignSupplySourceAppstoreSearchTab CampaignSupplySource = "APPSTORE_SEARCH_TAB"
	// CampaignSupplySourceAppstoreTodayTab is for a campaign supply source on APPSTORE_TODAY_TAB.
	CampaignSupplySourceAppstoreTodayTab CampaignSupplySource = "APPSTORE_TODAY_TAB"
	// CampaignSupplySourceAppstoreProductPagesBrowse is for a campaign supply source on APPSTORE_PRODUCT_PAGES_BROWSE.
	CampaignSupplySourceAppstoreProductPagesBrowse CampaignSupplySource = "APPSTORE_PRODUCT_PAGES_BROWSE"
)

// CampaignServingStatus is the status of the campaign.
type CampaignServingStatus string

const (
	// CampaignServingStatusRunning is for a campaign serving status source on RUNNING.
	CampaignServingStatusRunning CampaignServingStatus = "RUNNING"
	// CampaignServingStatusNotRunning is for a campaign supply source on NOT_RUNNING.
	CampaignServingStatusNotRunning CampaignServingStatus = "NOT_RUNNING"
)

// CampaignStatus is the user-controlled status to enable or pause the campaign.
type CampaignStatus string

const (
	// CampaignStatusEnabled is for a campaign status on ENABLED.
	CampaignStatusEnabled CampaignStatus = "ENABLED"
	// CampaignStatusPaused is for a campaign status source on PAUSED.
	CampaignStatusPaused CampaignStatus = "PAUSED"
)

// CampaignCountryOrRegionServingStateReasons is the reasons why a campaign can’t run
//
// https://developer.apple.com/documentation/apple_search_ads/campaign/countryorregionservingstatereasons
type CampaignCountryOrRegionServingStateReasons map[string]CampaignCountryOrRegionServingStateReason

// CampaignCountryOrRegionServingStateReason is a reason that returns when a campaign can’t run for a specified country or region.
type CampaignCountryOrRegionServingStateReason string

const (
	// CampaignCountryOrRegionServingStateReasonAppNotEligible is for a campaign country or region serving state reason on APP_NOT_ELIGIBLE.
	CampaignCountryOrRegionServingStateReasonAppNotEligible CampaignCountryOrRegionServingStateReason = "APP_NOT_ELIGIBLE"
	// CampaignCountryOrRegionServingStateReasonAppNotEligibleSearchAds is for a campaign country or region serving state reason on APP_NOT_ELIGIBLE_SEARCHADS.
	CampaignCountryOrRegionServingStateReasonAppNotEligibleSearchAds CampaignCountryOrRegionServingStateReason = "APP_NOT_ELIGIBLE_SEARCHADS"
	// CampaignCountryOrRegionServingStateReasonAppNotPublishedYet is for a campaign country or region serving state reason on APP_NOT_PUBLISHED_YET.
	CampaignCountryOrRegionServingStateReasonAppNotPublishedYet CampaignCountryOrRegionServingStateReason = "APP_NOT_PUBLISHED_YET"
	// CampaignCountryOrRegionServingStateReasonSapinLawAgentUnknown is for a campaign country or region serving state reason on SAPIN_LAW_AGENT_UNKNOWN.
	CampaignCountryOrRegionServingStateReasonSapinLawAgentUnknown CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_AGENT_UNKNOWN"
	// CampaignCountryOrRegionServingStateReasonSapinLawFrenchBizUnknown is for a campaign country or region serving state reason on SAPIN_LAW_FRENCH_BIZ_UNKNOWN.
	CampaignCountryOrRegionServingStateReasonSapinLawFrenchBizUnknown CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_FRENCH_BIZ_UNKNOWN"
	// CampaignCountryOrRegionServingStateReasonSapinLawFrenchBiz is for a campaign country or region serving state reason on SAPIN_LAW_FRENCH_BIZ.
	CampaignCountryOrRegionServingStateReasonSapinLawFrenchBiz CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_FRENCH_BIZ"
)

// Campaign is the response to a request to create and fetch campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/campaign
type Campaign struct {
	AdamID                             int64                                      `json:"adamId,omitempty"`
	AdChannelType                      CampaignAdChannelType                      `json:"adChannelType,omitempty"`
	BillingEvent                       BillingEventType                           `json:"billingEvent,omitempty"`
	BudgetAmount                       *Money                                     `json:"budgetAmount,omitempty"`
	BudgetOrders                       []int64                                    `json:"budgetOrders,omitempty"`
	CountriesOrRegions                 []string                                   `json:"countriesOrRegions,omitempty"`
	CountryOrRegionServingStateReasons CampaignCountryOrRegionServingStateReasons `json:"countryOrRegionServingStateReasons,omitempty"`
	CreationTime                       DateTime                                   `json:"creationTime,omitempty"`
	DailyBudgetAmount                  *Money                                     `json:"dailyBudgetAmount,omitempty"`
	Deleted                            bool                                       `json:"deleted,omitempty"`
	DisplayStatus                      CampaignDisplayStatus                      `json:"displayStatus,omitempty"`
	EndTime                            *DateTime                                  `json:"endTime,omitempty"`
	ID                                 int64                                      `json:"id,omitempty"`
	LocInvoiceDetails                  *LOCInvoiceDetails                         `json:"locInvoiceDetails,omitempty"`
	ModificationTime                   DateTime                                   `json:"modificationTime,omitempty"`
	Name                               string                                     `json:"name,omitempty"`
	OrgID                              int64                                      `json:"orgId,omitempty"`
	PaymentModel                       PaymentModel                               `json:"paymentModel,omitempty"`
	ServingStateReasons                []CampaignServingStateReason               `json:"servingStateReasons,omitempty"`
	ServingStatus                      CampaignServingStatus                      `json:"servingStatus,omitempty"`
	StartTime                          DateTime                                   `json:"startTime,omitempty"`
	Status                             CampaignStatus                             `json:"status,omitempty"`
	SupplySources                      []CampaignSupplySource                     `json:"supplySources,omitempty"`
}

// LOCInvoiceDetails is the response to a request to fetch campaign details for a standard invoicing payment model
//
// https://developer.apple.com/documentation/apple_search_ads/locinvoicedetails
type LOCInvoiceDetails struct {
	BillingContactEmail string `json:"billingContactEmail,omitempty"`
	BuyerEmail          string `json:"buyerEmail,omitempty"`
	BuyerName           string `json:"buyerName,omitempty"`
	ClientName          string `json:"clientName,omitempty"`
	OrderNumber         string `json:"orderNumber,omitempty"`
}

// GetAllCampaignQuery defines query parameter for GetAllCampaigns endpoint.
type GetAllCampaignQuery struct {
	Limit  int32 `form:"limit" param:"default=10"`
	Offset int32 `form:"offset"`
}

// CampaignUpdate is the list of campaign fields that are updatable
//
// https://developer.apple.com/documentation/apple_search_ads/campaignupdate
type CampaignUpdate struct {
	BudgetAmount       *Money             `json:"budgetAmount,omitempty"`
	BudgetOrders       int64              `json:"budgetOrders,omitempty"`
	CountriesOrRegions []string           `json:"countriesOrRegions,omitempty"`
	DailyBudgetAmount  *Money             `json:"dailyBudgetAmount,omitempty"`
	LOCInvoiceDetails  *LOCInvoiceDetails `json:"locInvoiceDetails,omitempty"`
	Name               string             `json:"name,omitempty"`
	Status             *CampaignStatus    `json:"status,omitempty"`
}

// UpdateCampaignRequest is the payload properties to clear Geo Targeting from a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/updatecampaignrequest
type UpdateCampaignRequest struct {
	Campaign                                 *CampaignUpdate `json:"campaign,omitempty"`
	ClearGeoTargetingOnCountryOrRegionChange bool            `json:"clearGeoTargetingOnCountryOrRegionChange,omitempty"`
}

// CampaignListResponse is the response details of campaign requests
//
// https://developer.apple.com/documentation/apple_search_ads/campaignlistresponse
type CampaignListResponse struct {
	Campaigns  []*Campaign        `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// CampaignResponse is a container for the campaign response body
//
// https://developer.apple.com/documentation/apple_search_ads/campaignresponse
type CampaignResponse struct {
	Campaign   *Campaign          `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// GetAllCampaigns Fetches all of an organization’s assigned campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_campaigns
func (s *CampaignService) GetAllCampaigns(params *GetAllCampaignQuery) (*CampaignListResponse, error) {
	url := "campaigns"
	res := new(CampaignListResponse)
	err := s.client.get(url, res, params)
	return res, err
}

// GetCampaign Fetches a specific campaign by campaign identifier
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_campaign
func (s *CampaignService) GetCampaign(campaignID int64) (*CampaignResponse, error) {
	url := fmt.Sprintf("campaigns/%d", campaignID)
	res := new(CampaignResponse)
	err := s.client.get(url, res)
	return res, err
}

// FindCampaigns Fetches campaigns with selector operators
//
// https://developer.apple.com/documentation/apple_search_ads/find_campaigns
func (s *CampaignService) FindCampaigns(selector *Selector) (*CampaignListResponse, error) {
	url := "campaigns/find"
	res := new(CampaignListResponse)
	err := s.client.post(url, res, selector)

	return res, err
}

// CreateCampaign Creates a campaign to promote an app
//
// https://developer.apple.com/documentation/apple_search_ads/create_a_campaign
func (s *CampaignService) CreateCampaign(campaign *Campaign) (*CampaignResponse, error) {
	url := "campaigns"
	res := new(CampaignResponse)
	err := s.client.post(url, res, campaign)

	return res, err
}

// DeleteCampaign Deletes a specific campaign by campaign identifier
//
// https://developer.apple.com/documentation/apple_search_ads/delete_a_campaign
func (s *CampaignService) DeleteCampaign(campaignID int64) (*BaseResponse, error) {
	url := fmt.Sprintf("campaigns/%d", campaignID)
	res := new(BaseResponse)
	err := s.client.delete(url, res)

	return res, err
}

// UpdateCampaign Updates a campaign with a campaign identifier
//
// https://developer.apple.com/documentation/apple_search_ads/update_a_campaign
func (s *CampaignService) UpdateCampaign(campaignID int64, req *UpdateCampaignRequest) (*CampaignResponse, error) {
	url := fmt.Sprintf("campaigns/%d", campaignID)
	res := new(CampaignResponse)
	err := s.client.put(url, res, req)

	return res, err
}
