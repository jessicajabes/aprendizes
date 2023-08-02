package dtos

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Secretary int    `json:"secretary"`
	Admin     bool   `json:"admin"`
}
