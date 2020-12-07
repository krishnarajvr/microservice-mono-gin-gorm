package common

// Response defines the api response
type Response struct {
	Status int    `json:"code" example:"200"`
	Data   string `json:"data" example:"{users}"`
	Error  string `json:"error" example:""`
}
