package asa

// AccessControlListService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/calling_the_apple_search_ads_api
type AccessControlListService service

// ErrorResponseItemMessageCode is a system-assigned error code.
type ErrorResponseItemMessageCode string

const (
	// ErrorResponseItemMessageCodeUnauthorized is for an error response item message code on UNAUTHORIZED.
	ErrorResponseItemMessageCodeUnauthorized ErrorResponseItemMessageCode = "UNAUTHORIZED"
	// ErrorResponseItemMessageCodeInvalidDateFormat is for an error response item message code on INVALID_DATE_FORMAT.
	ErrorResponseItemMessageCodeInvalidDateFormat ErrorResponseItemMessageCode = "INVALID_DATE_FORMAT"
)

// PaymentModel is the payment model that you set through the Search Ads UI.
type PaymentModel string

const (
	// PaymentModelPayG is a pay-as-you-go payment mode.
	PaymentModelPayG PaymentModel = "PAYG"
	// PaymentModelLoc is a line-of-credit payment model.
	PaymentModelLoc PaymentModel = "LOC"
)

// UserACLRoleName governs what a user can see and do within the account.
type UserACLRoleName string

const (
	// UserACLRoleNameAPIAccountManager is for Manage all campaigns within an account with read-and-write capabilities.
	UserACLRoleNameAPIAccountManager UserACLRoleName = "API Account Manager"
	// UserACLRoleNameAPIAccountReadOnly is for View reporting across the account with read-only permission.
	UserACLRoleNameAPIAccountReadOnly UserACLRoleName = "API Account Read Only"
	// UserACLRoleNameLimitedAccessAPIReadWrite is for View reporting.
	UserACLRoleNameLimitedAccessAPIReadWrite UserACLRoleName = "Limited Access: API Read & Write"
	// UserACLRoleNameLimitedAccessAPIReadOnly is View reporting across the organization.
	UserACLRoleNameLimitedAccessAPIReadOnly UserACLRoleName = "Limited Access: API Read Only"
)

// ReportingRequestTimeZone is the default timeZone during account creation through the Apple Search Ads UI.
type ReportingRequestTimeZone string

const (
	// ReportingRequestTimeZoneUTC is for a reporting request timezone on UTC.
	ReportingRequestTimeZoneUTC ReportingRequestTimeZone = "UTC"
	// ReportingRequestTimeZoneORTZ is for a reporting request timezone on ORTZ (organization time zone).
	ReportingRequestTimeZoneORTZ ReportingRequestTimeZone = "ORTZ"
)

// UserACL is the response to ACL requests
//
// https://developer.apple.com/documentation/apple_search_ads/useracl
type UserACL struct {
	Currency     string                   `json:"currency,omitempty"`
	DisplayName  string                   `json:"displayName,omitempty"`
	OrgID        int64                    `json:"orgId,omitempty"`
	OrgName      string                   `json:"orgName,omitempty"`
	ParentOrgID  int64                    `json:"parentOrgId,omitempty"`
	PaymentModel PaymentModel             `json:"paymentModel,omitempty"`
	RoleNames    []UserACLRoleName        `json:"roleNames,omitempty"`
	TimeZone     ReportingRequestTimeZone `json:"timeZone,omitempty"`
}

// ErrorResponseItem is the error response details in the response body
//
// https://developer.apple.com/documentation/apple_search_ads/errorresponseitem
type ErrorResponseItem struct {
	Field       string                       `json:"field,omitempty"`
	Message     string                       `json:"message,omitempty"`
	MessageCode ErrorResponseItemMessageCode `json:"messageCode,omitempty"`
}

// ErrorResponseBody is a container for the error response body
//
// https://developer.apple.com/documentation/apple_search_ads/errorresponsebody
type ErrorResponseBody struct {
	Errors []ErrorResponseItem `json:"errors,omitempty"`
}

// PageDetail is the number of items that return in the page
//
// https://developer.apple.com/documentation/apple_search_ads/pagedetail
type PageDetail struct {
	TotalResults int `json:"totalResults,omitempty"`
	StartIndex   int `json:"startIndex,omitempty"`
	ItemsPerPage int `json:"itemsPerPage,omitempty"`
}

// UserACLListResponse is a container for ACL call responses
//
// https://developer.apple.com/documentation/apple_search_ads/useracllistresponse
type UserACLListResponse struct {
	UserAcls   []*UserACL         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// APIErrorResponse A container for the error response body
//
// https://developer.apple.com/documentation/apple_search_ads/apierrorresponse
type APIErrorResponse struct {
	Error ErrorResponseBody `json:"error,omitempty"`
}

type BaseResponse struct {
	Data       interface{}        `json:"data,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
}

// GetUserACL Fetches roles and organizations that the API has access to
//
// https://developer.apple.com/documentation/apple_search_ads/get_user_acl
func (s *AccessControlListService) GetUserACL() (*UserACLListResponse, error) {
	url := "acls"
	resp := new(UserACLListResponse)
	err := s.client.get(url, resp)
	return resp, err
}
