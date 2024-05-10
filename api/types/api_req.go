package telegraph

type ApiReq struct {
	// Endpoint to which we address our API request
	Endpoint string

	// Params of the API request
	Params interface{}

	// Model for the API request
	Model interface{}
}
