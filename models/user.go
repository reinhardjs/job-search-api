package models

type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Role     string `json:"role,omitempty"`
}
