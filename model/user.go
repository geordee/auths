package model

// User model
type User struct {
	ID     string   `json:"id"`
	Orgs   []string `json:"orgs"`
	Roles  []string `json:"roles"`
	Scopes []string `json:"scopes"`
}
