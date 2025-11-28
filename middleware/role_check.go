package middleware

import (
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func CheckRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(shared.KeyRequester).(shared.Requester)

		for _, role := range roles {
			if role == requester.GetRole() {
				c.Next()
				return
			}
		}

		c.Abort()

		panic(shared.ErrUnauthorized.WithError("user role not allowed"))
	}
}