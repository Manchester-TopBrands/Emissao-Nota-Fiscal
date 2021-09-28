package models

// Login ...
type Login struct {
	Username     string `json:"username,omitempty"`
	Userpassword string `json:"userpassword,omitempty"`
}
type Logout struct {
	Token string `json:"token,omitempty"`
}
