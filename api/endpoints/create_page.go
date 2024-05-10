package telegraph

import (
	telegraph "go-telegraph/api/types"
	"go-telegraph/pkg"
)

// This object represents a Telegraph create page request params.
type CreatePageParams struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token" validate:"nonnil"`

	// Required. Page title.
	Title string `json:"title" validate:"nonnil,min=1,max=256"`

	// Author name, displayed below the article's title.
	AuthorName string `json:"author_name,omitempty" validate:"min=0,max=128"`

	// Profile link, opened when users click on the author's name below the title.
	// Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url,omitempty" validate:"min=0,max=512"`

	// Required. Content of the page.
	// Array of Node, up to 64 KB
	Content []telegraph.Node `json:"content" validate:"nonnil"`

	// If true, a content field will be returned in the Page object.
	// false by default
	ReturnContent bool `json:"return_content" validate:"nonnil"`
}

// Use this method to create a new Telegraph page.
// On success, returns a Page object.
func (c *Client) CreatePage(params *CreatePageParams) (page *telegraph.Page, err error) {
	// If token is not passed, try to fetch latest from client
	if len(params.AccessToken) == 0 {
		params.AccessToken = c.GetAccessToken()
	}

	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "createPage",
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
