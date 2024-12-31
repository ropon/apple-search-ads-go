package asa

// GeoService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/search_apps_and_geolocations
type GeoService service

// GeoEntityType is the locations available for targeting.
type GeoEntityType string

const (
	// GeoEntityTypeCountry is for a geo targeting locations on Country.
	GeoEntityTypeCountry GeoEntityType = "Country"
	// GeoEntityTypeAdminArea is for a geo targeting locations on AdminArea.
	GeoEntityTypeAdminArea GeoEntityType = "AdminArea"
	// GeoEntityTypeLocality is for a geo targeting locations on Locality.
	GeoEntityTypeLocality GeoEntityType = "Locality"
)

// SearchGeoQuery defines query parameter for SearchGeos endpoint.
type SearchGeoQuery struct {
	Limit       int32         `form:"limit,omitempty"`
	Offset      int32         `form:"offset,omitempty"`
	Query       string        `form:"query,omitempty"`
	CountryCode string        `form:"countrycode,omitempty"`
	Entity      GeoEntityType `form:"entity,omitempty"`
}

// ListGeoQuery defines query parameter for GetGeos endpoint.
type ListGeoQuery struct {
	Limit  int32 `form:"limit,omitempty"`
	Offset int32 `form:"offset,omitempty"`
}

// GeoRequest is the geosearch request
//
// https://developer.apple.com/documentation/apple_search_ads/georequest
type GeoRequest struct {
	Entity GeoEntityType `json:"entity"`
	ID     string        `json:"id"`
}

// SearchEntityListResponse is the response details of geosearch requests
//
// https://developer.apple.com/documentation/apple_search_ads/searchentitylistresponse
type SearchEntityListResponse struct {
	SearchEntities []*SearchEntity    `json:"data,omitempty"`
	Error          *ErrorResponseBody `json:"error,omitempty"`
	Pagination     *PageDetail        `json:"pagination,omitempty"`
}

// SearchEntity is the list of geolocations that includes the geoidentifier and entity type
//
// https://developer.apple.com/documentation/apple_search_ads/searchentity
type SearchEntity struct {
	DisplayName string `json:"displayName,omitempty"`
	Entity      string `json:"entity,omitempty"`
	ID          string `json:"id,omitempty"`
}

// SearchGeos Fetches a list of geolocations for audience refinement
//
// https://developer.apple.com/documentation/apple_search_ads/search_for_geolocations
func (s *GeoService) SearchGeos(params *SearchGeoQuery) (*SearchEntityListResponse, error) {
	url := "search/geo"
	res := new(SearchEntityListResponse)
	err := s.client.get(url, res, params)

	return res, err
}

// GetGeos Gets geolocation details using a geoidentifier
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_list_of_geolocations
func (s *GeoService) GetGeos(query *ListGeoQuery, params []*GeoRequest) (*SearchEntityListResponse, error) {
	url := "search/geo"
	res := new(SearchEntityListResponse)
	err := s.client.postWithQuery(url, res, query, params)

	return res, err
}
