package middleware

import (
	"fmt"
	"log"
	"os"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

type CanGetStatusCode interface {
	StatusCode() int
}

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			isProduction := os.Getenv("ENV") == "prod" || os.Getenv("GIN_MODE") == "release"

			if r := recover(); r != nil {
				if appError, ok := r.(CanGetStatusCode); ok {
					c.JSON(appError.StatusCode(), appError)

					if !isProduction {
						log.Printf("Error: %+v", appError)
						// panic(r)
					}
					return
				}

				appError := shared.ErrInternalServerError.WithDebug(fmt.Sprintf("%s", r))

				if isProduction {
					appError.WithDebug("")
				}

				c.JSON(appError.StatusCode(), appError)

				if !isProduction {
					panic(r)
				}
			}
		}()

		c.Next()
	}
}