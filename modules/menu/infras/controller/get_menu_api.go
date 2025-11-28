package menucontroller

import (
	"net/http"
	menuservice "vht-go/modules/menu/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPMenuController) GetMenuByIdAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid menu id"})
			return
		}

		menu, err := ctrl.getHandler.Handle(c.Request.Context(), &menuservice.GetMenuQuery{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": menu})
	}
}

