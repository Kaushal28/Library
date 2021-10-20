package entities

// standard response format
type Response struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}
