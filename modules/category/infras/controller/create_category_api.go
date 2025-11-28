package categorycontroller

import (
	"net/http"
	categorydtos "vht-go/modules/category/dtos"
	categoryservice "vht-go/modules/category/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPCategoryController) CreateCategoryAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto categorydtos.CreateCategoryDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := ctrl.createHandler.Handle(c.Request.Context(), &categoryservice.CreateCategoryResultCommand{DTO: &dto})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": newId,
		})
	}
}