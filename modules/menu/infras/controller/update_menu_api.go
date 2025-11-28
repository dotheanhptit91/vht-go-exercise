package menucontroller

import (
	"net/http"
	menudtos "vht-go/modules/menu/dtos"
	menuservice "vht-go/modules/menu/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPMenuController) UpdateMenuAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid menu id"})
			return
		}

		var dto menudtos.UpdateMenuDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = ctrl.updateHandler.Handle(c.Request.Context(), &menuservice.UpdateMenuCommand{
			Id:  id,
			DTO: &dto,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

