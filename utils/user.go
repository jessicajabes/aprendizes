package utils

import "context"

type userKey struct{}
type adminKey struct{}

func ContextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userKey{}, userID)
}

func ContextWithUserAdmin(ctx context.Context, Admin bool) context.Context {
	return context.WithValue(ctx, adminKey{}, Admin)
}

func GetUserIDFromContext(ctx context.Context) (int, bool) {
	value, exists := ctx.Value(userKey{}).(int)
	return value, exists
}

func GetUserAdminFromContext(ctx context.Context) (bool, bool) {
	value, exists := ctx.Value(adminKey{}).(bool)
	return value, exists
}
