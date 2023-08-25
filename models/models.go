package models

// User schema of the user table
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Email_is string `json:"email_is,omitempty"`
	Token    string `json:"token"`
	Location string `json:"location"`
}
