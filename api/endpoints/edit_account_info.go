package telegraph

import (
	telegraph "go-telegraph/api/types"
	"go-telegraph/pkg"
)

// This object represents a edit Telegraph account info request params.
type EditAccountInfoParams struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token" validate:"nonnil"`

	// New account name.
	ShortName string `json:"short_name" validate:"min=1,max=32"`

	// New default author name used when creating new articles.
	AuthorName string `json:"author_name,omitempty" validate:"min=0,max=128"`

	// New default profile link, opened when users click on the author's name below the title.
	// Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url,omitempty" validate:"min=0,max=512"`
}

// Use this method to update information about a Telegraph account.
// Pass only the parameters that you want to edit.
// On success, returns an Account object with the default fields.
func (c *Client) EditAccountInfo(params *EditAccountInfoParams) (account *telegraph.Account, err error) {
	// If token is not passed, try to fetch latest from client
	if len(params.AccessToken) == 0 {
		params.AccessToken = c.GetAccessToken()
	}

	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "editAccountInfo",
		Params:   params,
		Model:    telegraph.Account{},
	}

	// Make request
	result, err := pkg.PerformRequest(c.URL, &apiReq)
	if err != nil {
		return
	}

	// Cast interface{} -> Model type
	account = result.(*telegraph.Account)

	return
}
