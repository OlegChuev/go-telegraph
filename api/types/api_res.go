package telegraph

type ApiRes struct {
	// If ok equals true, the request was successful,
	// and the result of the query can be found in the result field.
	OK bool `json:"ok"`

	// Result of the API query execution
	Result interface{} `json:"result"`

	// In case of an unsuccessful request, ok equals false,
	// and the error is explained in the error field
	Error string `json:"error"`
}
