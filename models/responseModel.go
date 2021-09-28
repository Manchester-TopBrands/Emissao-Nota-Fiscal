package models

// Response ...
type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// RespLink ...
type RespLink struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
	PDF    string `json:"pdf,omitempty"`
	XML    string `json:"xml,omitempty"`
}
