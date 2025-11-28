package categorycontroller

import (
	"net/http"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"
	categoryservice "vht-go/modules/category/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

type ListCategoriesRequest struct {
	shared.Paging `form:"paging"`
	categorydtos.FilterStatusDTO `form:"filter"`
}

func (ctrl *HTTPCategoryController) GetListCategoriesAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData ListCategoriesRequest
		if err := c.ShouldBindQuery(&reqData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		reqData.Paging.Process()

		dto := categorydtos.ListCategoryDTO{
			Paging: &reqData.Paging,
			Filter: &reqData.FilterStatusDTO,
		}

		var categories []categorydomain.Category
		categories, err := ctrl.listHandler.Handle(c.Request.Context(), &categoryservice.ListCategoryQuery{DTO: &dto})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": categories,
			"paging": reqData.Paging,
		})
	}
}