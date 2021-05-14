package models

//User struct declaration
type User struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}
