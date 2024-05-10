package telegraph

import (
	telegraph "github.com/OlegChuev/go-telegraph/api/types"
	"github.com/OlegChuev/go-telegraph/pkg"
)

type GetPageParams struct {
	// Required. Path to the Telegraph page
	// (in the format Title-12-31, i.e. everything that comes after http://telegra.ph/).
	Path string `json:"path" validate:"nonnil"`

	// If true, content field will be returned in Page object.
	// false by default
	ReturnContent bool `json:"return_content" validate:"nonnil"`
}

// Use this method to get a Telegraph page.
// Returns a Page object on success.
func (c *Client) GetPage(params *GetPageParams) (page *telegraph.Page, err error) {
	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "getPage",
		Params:   params,
		Model:    telegraph.Page{},
	}

	// Make request
	result, err := pkg.PerformRequest(c.URL, &apiReq)
	if err != nil {
		return
	}

	// Cast interface{} -> Model type
	page = result.(*telegraph.Page)

	return
}
