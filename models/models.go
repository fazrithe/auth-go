package models

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Email_is string `json:"email_is"`
	Token    string `json:"token"`
	Location string `json:"location"`
}
