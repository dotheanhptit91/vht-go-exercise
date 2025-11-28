package usercontroller

import (
	"net/http"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPUserController) GetMeAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(shared.KeyRequester).(shared.Requester)
		// requester.Subject() // user id

		c.JSON(http.StatusOK, shared.SimpleResponse(requester))
	}
}