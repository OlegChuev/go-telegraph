package telegraph

import (
	telegraph "go-telegraph/api/types"
	"go-telegraph/pkg"
)

// This object represents a get Telegraph account info request params.
type GetViewsParams struct {
	// Required. Path to the Telegraph page (in the format Title-12-31, where 12
	// is the month and 31 the day the article was first published).
	Path string `json:"path" validate:"nonnil"`

	// Required if month is passed.
	// If passed, the number of page views for the requested year will be returned.
	// Minimum value is 2000
	// Maximum value is 2100
	Year int `json:"year,omitempty"`

	// Required if day is passed.
	// If passed, the number of page views for the requested month will be returned.
	// Minimum value is 1
	// Maximum value is 12
	Month int `json:"month,omitempty"`

	// Required if hour is passed.
	// If passed, the number of page views for the requested day will be returned.
	// Minimum value is 1
	// Maximum value is 31
	Day int `json:"day,omitempty"`

	// If passed, the number of page views for the requested hour will be returned.
	// Minimum value is 0
	// Maximum value is 24
	Hour int `json:"hour,omitempty"`
}

func (c *Client) GetViews(params *GetViewsParams) (pageViews *telegraph.PageViews, err error) {
	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "getViews",
		Params:   params,
		Model:    telegraph.PageViews{},
	}

	// Make request
	result, err := pkg.PerformRequest(c.URL, &apiReq)
	if err != nil {
		return
	}

	// Cast interface{} -> Model type
	pageViews = result.(*telegraph.PageViews)

	return
}
