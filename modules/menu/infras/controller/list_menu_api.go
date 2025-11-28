package menucontroller

import (
	"net/http"
	menudtos "vht-go/modules/menu/dtos"
	menuservice "vht-go/modules/menu/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPMenuController) ListMenuAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto menudtos.ListMenuDTO

		if err := c.ShouldBindQuery(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := ctrl.listHandler.Handle(c.Request.Context(), &menuservice.ListMenuQuery{DTO: &dto})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

