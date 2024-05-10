package telegraph

import (
	telegraph "github.com/OlegChuev/go-telegraph/api/types"
	"github.com/OlegChuev/go-telegraph/pkg"
)

// This object represents a get Telegraph page list request params.
type GetPageListParams struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token" validate:"nonnil"`

	// Sequential number of the first page to be returned.
	// default = 0
	Offset int `json:"offset,omitempty"`

	// Limits the number of pages to be retrieved.
	// default = 50
	Limit int `json:"limit,omitempty" validate:"min=0,max=200"`
}

// Use this method to get a list of pages belonging to a Telegraph account.
// Returns a PageList object, sorted by most recently created pages first.
func (c *Client) GetPageList(params *GetPageListParams) (pageList *telegraph.PageList, err error) {
	// If token is not passed, try to fetch latest from client
	if len(params.AccessToken) == 0 {
		params.AccessToken = c.GetAccessToken()
	}

	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "getPageList",
		Params:   params,
		Model:    telegraph.PageList{},
	}

	// Make request
	result, err := pkg.PerformRequest(c.URL, &apiReq)
	if err != nil {
		return
	}

	// Cast interface{} -> Model type
	pageList = result.(*telegraph.PageList)

	return
}
