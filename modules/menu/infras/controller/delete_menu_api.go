package menucontroller

import (
	"net/http"
	menuservice "vht-go/modules/menu/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPMenuController) DeleteMenuAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid menu id"})
			return
		}

		err = ctrl.deleteHandler.Handle(c.Request.Context(), &menuservice.DeleteMenuCommand{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

