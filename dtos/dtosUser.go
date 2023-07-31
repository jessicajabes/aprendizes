package dtos

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"nomeusuario"`
	Password  string `json:"senha"`
	Secretary string `json:"secretaria"`
}
