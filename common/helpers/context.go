package helpers

import (
	"context"
	"strconv"
)

func GetJWTFromContext(ctx context.Context) string {
	jwt, ok := ctx.Value("jwt").(string)
	if !ok {
		return ""
	}

	return jwt
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return 0
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return 0
	}

	return id
}

func GetUserIDFromContextStr(ctx context.Context) string {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return ""
	}

	return userID
}

func GetExpFromContext(ctx context.Context) string {
	exp, ok := ctx.Value("jwtExp").(int64)
	if !ok {
		return ""
	}

	return strconv.FormatInt(exp, 10)
}

func GetRoleFromContext(ctx context.Context) string {
	role, ok := ctx.Value("role").(string)
	if !ok {
		return ""
	}

	return role
}

func GetPlanExpirationFromContext(ctx context.Context) int64 {
	exp, ok := ctx.Value("planExpirationDate").(int64)
	if !ok {
		return 0
	}

	return exp
}

func SetStatusCodeOnContext(ctx context.Context, statusCode int) context.Context {
	return context.WithValue(ctx, "statusCode", statusCode)
}
