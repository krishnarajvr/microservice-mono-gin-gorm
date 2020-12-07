package common

// Error500 defines the response error
type Error500 struct {
	Status int    `json:"code" example:"500"`
	Data   string `json:"data" example:""`
	Error  string `json:"error" example:"Internal Server Error"`
}
