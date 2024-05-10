package telegraph

import telegraph "github.com/OlegChuev/go-telegraph/api/types"

// IClient defines the interface for the Telegraph API client.
type IClient interface {
	// General methods

	// SetAccessToken sets the access token for the client.
	SetAccessToken(string)
	// GetAccessToken retrieves the access token of the client.
	GetAccessToken() string
	// SetURL sets the API url for the client.
	SetURL(string)

	// API Account related methods
	CreateAccount(*CreateAccountParams) (*telegraph.Account, error)
	EditAccountInfo(*EditAccountInfoParams) (*telegraph.Account, error)
	GetAccountInfo(*GetAccountInfoParams) (*telegraph.Account, error)
	RevokeAccessToken(*RevokeAccessTokenParams) (*telegraph.Account, error)

	// API Page related methods
	CreatePage(*CreatePageParams) (*telegraph.Page, error)
	EditPage(*EditPageParams) (*telegraph.Page, error)
	GetPage(*GetPageParams) (*telegraph.Page, error)
	GetViews(*GetViewsParams) (*telegraph.PageViews, error)
	GetPageList(*GetPageListParams) (*telegraph.PageList, error)
}

// Client implements the IClient interface and provides methods to interact with the Telegraph API.
type Client struct {
	// Can be either https://api.telegra.ph/ (default) OR https://api.graph.org (mirror)
	URL string

	// AccessToken
	AccessToken string
}

func (c *Client) SetAccessToken(token string) {
	c.AccessToken = token
}

func (c *Client) GetAccessToken() string {
	return c.AccessToken
}

func (c *Client) SetURL(url string) {
	c.URL = url
}

// NewClient creates a new instance of the Telegraph API client with default URL.
// URL can be specified later if needed.
func NewClient() IClient {
	return &Client{
		// Base telegra.ph API Endpoint
		URL: "https://api.telegra.ph",
	}
}
