package dtos

type Usuario struct {
	ID          int    `json:"id"`
	NomeUsuario string `json:"nomeusuario"`
	Senha       string `json:"senha"`
	Secretaria  string `json:"secretaria"`
}
