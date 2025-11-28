package categorycontroller

import (
	"net/http"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"
	categoryservice "vht-go/modules/category/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *HTTPCategoryController) GetCategoryByIdAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
	
		if err != nil {
			panic(shared.ErrBadRequest.WithWrap(err).WithDebug(err.Error()))
		}

		var category *categorydomain.Category
		dtos := &categorydtos.GetCategoryDTO{Id: &id}
		cmd := &categoryservice.GetCategoryQuery{DTO: dtos}

		category, err = ctrl.getHandler.Handle(c.Request.Context(), cmd)



		if err != nil {
			if err.Error() == categorydomain.ErrCategoryNotFound {
				c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, category)
	}
}