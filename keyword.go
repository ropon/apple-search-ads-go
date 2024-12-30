package asa

import (
	"fmt"
)

// KeywordService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/targeting_keywords_and_negative_keywords
type KeywordService service

// KeywordMatchType is an automated keyword and bidding strategy.
type KeywordMatchType string

const (
	// KeywordMatchTypeBroad is used this value to ensure your ads don’t run on relevant, close variants of a keyword, such as singulars, plurals, misspellings, synonyms, related searches, and phrases that include that term (fully or partially).
	KeywordMatchTypeBroad KeywordMatchType = "Broad"
	// KeywordMatchTypeExact is used this value for the most control over searches your ad may appear in. You can target a specific term and its close variants, such as common misspellings and plurals. Your ad may receive fewer impressions as a result, but your tap-through rates (TTRs) and conversions on those impressions may be higher because you’re reaching users most interested in your app.
	KeywordMatchTypeExact KeywordMatchType = "Exact"
)

// KeywordStatus defines model for Keyword Status.
type KeywordStatus string

const (
	// KeywordStatusActive is for a keyword status on Active state.
	KeywordStatusActive KeywordStatus = "ACTIVE"
	// KeywordStatusPaused is for a keyword status on Paused state.
	KeywordStatusPaused KeywordStatus = "PAUSED"
)

// Keyword defines model for Keyword.
//
// https://developer.apple.com/documentation/apple_search_ads/keyword
type Keyword struct {
	AdGroupID        int64            `json:"adGroupId,omitempty"`
	BidAmount        Money            `json:"bidAmount,omitempty"`
	CampaignID       int64            `json:"campaignID,omitempty"`
	CreationTime     DateTime         `json:"creationTime,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	ID               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType,omitempty"`
	ModificationTime DateTime         `json:"modificationTime,omitempty"`
	Status           KeywordStatus    `json:"status,omitempty"`
	Text             string           `json:"text,omitempty"`
}

// NegativeKeyword Negative keyword parameters to use in requests and responses
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeyword
type NegativeKeyword struct {
	AdGroupID        int64            `json:"adGroupId,omitempty"`
	CampaignID       int64            `json:"campaignId,omitempty"`
	CreationTime     DateTime         `json:"creationTime,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	ID               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType,omitempty"`
	ModificationTime DateTime         `json:"modificationTime,omitempty"`
	Status           KeywordStatus    `json:"status,omitempty"`
	Text             string           `json:"text,omitempty"`
}

// GetAllTargetingKeywordsQuery defines query parameter for GetAllTargetingKeywords endpoint.
type GetAllTargetingKeywordsQuery struct {
	Limit  int32 `form:"limit,omitempty"`
	Offset int32 `form:"offset,omitempty"`
}

// GetAllNegativeKeywordsQuery defines query parameter for GetAllNegativeKeywords endpoint.
type GetAllNegativeKeywordsQuery struct {
	Limit  int32 `form:"limit,omitempty"`
	Offset int32 `form:"offset,omitempty"`
}

// KeywordUpdateRequest Targeting keyword parameters to use in requests and responses
//
// https://developer.apple.com/documentation/apple_search_ads/keywordupdaterequest
type KeywordUpdateRequest struct {
	AdGroupID        int64            `json:"adGroupId,omitempty"`
	BidAmount        *Money           `json:"bidAmount,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	ID               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType"`
	ModificationTime DateTime         `json:"modificationTime"`
}

// KeywordListResponse defines model for Keyword List Response.
//
// https://developer.apple.com/documentation/apple_search_ads/keywordlistresponse
type KeywordListResponse struct {
	Keywords   []*Keyword         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// KeywordResponse is a container for the targeting keywords response body.
//
// https://developer.apple.com/documentation/apple_search_ads/keywordresponse
type KeywordResponse struct {
	Keyword    *Keyword           `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// NegativeKeywordListResponse The response details of negative keyword requests
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeywordlistresponse
type NegativeKeywordListResponse struct {
	Keywords   []*NegativeKeyword `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// NegativeKeywordResponse is a container for the negative keyword response body
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeywordresponse
type NegativeKeywordResponse struct {
	NegativeKeyword *NegativeKeyword   `json:"data,omitempty"`
	Error           *ErrorResponseBody `json:"error,omitempty"`
	Pagination      *PageDetail        `json:"pagination,omitempty"`
}

// IntegerResponse is a common integer type response
//
// https://developer.apple.com/documentation/apple_search_ads/integerresponse
type IntegerResponse struct {
	Data       int32              `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// FindTargetingKeywords Fetches targeting keywords in a campaign’s ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/create_targeting_keywords
func (s *KeywordService) FindTargetingKeywords(campaignID int64, selector *Selector) (*KeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/targetingkeywords/find", campaignID)
	res := new(KeywordListResponse)
	err := s.client.post(url, res, selector)

	return res, err
}

// GetTargetingKeyword Fetches a specific targeting keyword in an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_targeting_keyword_in_an_ad_group
func (s *KeywordService) GetTargetingKeyword(campaignID int64, adGroupID int64, keywordID int64) (*KeywordResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/%d", campaignID, adGroupID, keywordID)
	res := new(KeywordResponse)
	err := s.client.get(url, res)

	return res, err
}

// GetAllTargetingKeywords Fetches all targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_targeting_keywords_in_an_ad_group
func (s *KeywordService) GetAllTargetingKeywords(campaignID int64, adGroupID int64, params *GetAllTargetingKeywordsQuery) (*KeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/", campaignID, adGroupID)
	res := new(KeywordListResponse)
	err := s.client.get(url, res, params)

	return res, err
}

// CreateTargetingKeywords Creates targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/create_targeting_keywords
func (s *KeywordService) CreateTargetingKeywords(campaignID int64, adGroupID int64, keyword []*Keyword) (*KeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignID, adGroupID)
	res := new(KeywordListResponse)
	err := s.client.post(url, res, keyword)

	return res, err
}

// UpdateTargetingKeywords Updates targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/update_targeting_keywords
func (s *KeywordService) UpdateTargetingKeywords(campaignID int64, adGroupID int64, updateRequests []*KeywordUpdateRequest) (*KeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignID, adGroupID)
	res := new(KeywordListResponse)
	err := s.client.put(url, res, updateRequests)

	return res, err
}

// FindNegativeKeywords Fetches negative keywords for campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/find_campaign_negative_keywords
func (s *KeywordService) FindNegativeKeywords(campaignID int64, selector *Selector) (*NegativeKeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/find", campaignID)
	res := new(NegativeKeywordListResponse)
	err := s.client.post(url, res, selector)

	return res, err
}

// FindAdGroupNegativeKeywords Fetches negative keywords in a campaign’s ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/find_ad_group_negative_keywords
func (s *KeywordService) FindAdGroupNegativeKeywords(campaignID int64, selector *Selector) (*NegativeKeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/negativekeywords/find", campaignID)
	res := new(NegativeKeywordListResponse)
	err := s.client.post(url, res, selector)

	return res, err
}

// GetNegativeKeyword Fetches a specific negative keyword in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_campaign_negative_keyword
func (s *KeywordService) GetNegativeKeyword(campaignID int64, keywordID int64) (*NegativeKeywordResponse, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/%d", campaignID, keywordID)
	res := new(NegativeKeywordResponse)
	err := s.client.get(url, res)

	return res, err
}

// GetAdGroupNegativeKeyword Fetches a specific negative keyword in an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/get_an_ad_group_negative_keyword
func (s *KeywordService) GetAdGroupNegativeKeyword(campaignID int64, adGroupID int64, keywordID int64) (*NegativeKeywordResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/%d", campaignID, adGroupID, keywordID)
	res := new(NegativeKeywordResponse)
	err := s.client.get(url, res)

	return res, err
}

// GetAllNegativeKeywords Fetches all negative keywords in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_campaign_negative_keywords
func (s *KeywordService) GetAllNegativeKeywords(campaignID int64, params *GetAllNegativeKeywordsQuery) (*NegativeKeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/", campaignID)
	res := new(NegativeKeywordListResponse)
	err := s.client.get(url, res, params)

	return res, err
}

// GetAllAdGroupNegativeKeywords Fetches all negative keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_ad_group_negative_keywords
func (s *KeywordService) GetAllAdGroupNegativeKeywords(campaignID int64, adGroupID int64, params *GetAllNegativeKeywordsQuery) (*NegativeKeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/", campaignID, adGroupID)
	res := new(NegativeKeywordListResponse)
	err := s.client.get(url, res, params)

	return res, err
}

// CreateNegativeKeywords Creates negative keywords for a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/create_campaign_negative_keywords
func (s *KeywordService) CreateNegativeKeywords(campaignID int64, keyword []*NegativeKeyword) (*NegativeKeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	res := new(NegativeKeywordListResponse)
	err := s.client.post(url, res, keyword)

	return res, err
}

// DeleteNegativeKeywords Deletes negative keywords from a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/delete_campaign_negative_keywords
func (s *KeywordService) DeleteNegativeKeywords(campaignID int64, keywordIds []int64) (*IntegerResponse, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/delete/bulk", campaignID)
	res := new(IntegerResponse)
	err := s.client.post(url, res, keywordIds)

	return res, err
}

// DeleteAdGroupNegativeKeywords Deletes negative keywords from an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/delete_ad_group_negative_keywords
func (s *KeywordService) DeleteAdGroupNegativeKeywords(campaignID int64, adGroupID int64, keywordIds []int64) (*IntegerResponse, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/delete/bulk", campaignID, adGroupID)
	res := new(IntegerResponse)
	err := s.client.post(url, res, keywordIds)

	return res, err
}

// UpdateNegativeKeywords Updates negative keywords in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/update_campaign_negative_keywords
func (s *KeywordService) UpdateNegativeKeywords(campaignID int64, updateRequests []*NegativeKeyword) (*NegativeKeywordListResponse, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	res := new(NegativeKeywordListResponse)
	err := s.client.put(url, res, updateRequests)

	return res, err
}
