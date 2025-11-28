package middleware

import (
	"errors"
	"strings"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func extractToken(authorization string) (string, error) {
	token := strings.TrimPrefix(authorization, "Bearer ")

	if token == "" {
		// return "", datatype.ErrUnauthorized.WithDebug("Token is required")
		return "", errors.New("token is required")
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		// return "", datatype.ErrUnauthorized.WithDebug("Token is required")
		return "", errors.New("token is required")
	}

	return token, nil
}

type ITokenIntrospector interface {
	Introspect(token string) (shared.Requester, error)
}

func Auth(tokenValidator ITokenIntrospector) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		requester, err := tokenValidator.Introspect(token)

		if err != nil {
			panic(shared.ErrUnauthorized.WithWrap(err).WithDebug(err.Error()))
		}

		c.Set(shared.KeyRequester, requester)
		c.Next()
	}
}