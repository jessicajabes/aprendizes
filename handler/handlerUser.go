package handler

type UserStorager interface {
}

type HandlerUser struct {
	service UserStorager
}

func NewHandlerUser(service UserStorager) *HandlerUser {
	return &HandlerUser{service: service}
}
