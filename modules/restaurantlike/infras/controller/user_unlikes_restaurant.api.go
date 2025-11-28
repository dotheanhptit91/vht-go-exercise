package rstlikecontroller

import (
	"net/http"
	"strconv"
	rstlikeservice "vht-go/modules/restaurantlike/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantLikeController) UserUnlikeRestaurantAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(shared.ErrBadRequest.WithError(err.Error()))
		}

		requester := c.MustGet(shared.KeyRequester).(shared.Requester)

		cmd := &rstlikeservice.UnlikeRestaurantCommand{
			RestaurantId: id,
			Requester:    requester,
		}

		if err := ctrl.unlikeRestaurantCommandHandler.Handle(c.Request.Context(), cmd); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, shared.SimpleResponse(true))
	}
}