package usercontroller

import (
	"net/http"
	userdto "vht-go/modules/user/dto"
	userservice "vht-go/modules/user/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPUserController) LoginUserAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userdto.LoginUserDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			panic(shared.ErrBadRequest.WithError(err.Error()))
		}

		cmd := userservice.LoginUserCommand{
			DTO: &dto,
		}

		response, err := ctrl.loginUserHandler.Handle(c.Request.Context(), &cmd)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, response)
	}
}