package categorycontroller

import (
	"net/http"
	categoryservice "vht-go/modules/category/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPCategoryController) DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = ctrl.deleteHandler.Handle(c.Request.Context(), &categoryservice.DeleteCategoryCommand{Id: &id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}