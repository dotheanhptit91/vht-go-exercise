package usercontroller

import (
	"net/http"
	userdto "vht-go/modules/user/dto"
	userservice "vht-go/modules/user/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPUserController) RegisterUserAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userdto.RegisterUserDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			panic(shared.ErrBadRequest.WithError(err.Error()))
		}

		cmd := userservice.RegisterUserCommand{
			DTO: &dto,
		}

		resp, err := ctrl.registerUserHandler.Handle(c.Request.Context(), &cmd)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, shared.SimpleResponse(resp))
	}
}