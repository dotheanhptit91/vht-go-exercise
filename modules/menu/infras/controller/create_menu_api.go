package menucontroller

import (
	"net/http"
	menudtos "vht-go/modules/menu/dtos"
	menuservice "vht-go/modules/menu/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPMenuController) CreateMenuAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto menudtos.CreateMenuDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := ctrl.createHandler.Handle(c.Request.Context(), &menuservice.CreateMenuResultCommand{DTO: &dto})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": newId})
	}
}

