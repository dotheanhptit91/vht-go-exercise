package usercontroller

import (
	"net/http"
	"vht-go/middleware"
	userdto "vht-go/modules/user/dto"
	userservice "vht-go/modules/user/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPUserController struct {
	registerUserHandler shared.ICommandResultHandler[*userservice.RegisterUserCommand, *uuid.UUID]
	loginUserHandler    shared.ICommandResultHandler[*userservice.LoginUserCommand, *userdto.LoginResponseDTO]
}

func NewHTTPUserController(
	registerUserHandler shared.ICommandResultHandler[*userservice.RegisterUserCommand, *uuid.UUID],
	loginUserHandler shared.ICommandResultHandler[*userservice.LoginUserCommand, *userdto.LoginResponseDTO],
) *HTTPUserController {
	return &HTTPUserController{
		registerUserHandler: registerUserHandler,
		loginUserHandler:    loginUserHandler,
	}
}

func (ctrl *HTTPUserController) SetupRouter(v1 *gin.RouterGroup, middlewareProvider middleware.IMiddlewareProvider) {
	v1.POST("/register", ctrl.RegisterUserAPI())
	v1.POST("/authenticate", ctrl.LoginUserAPI())

	v1.GET("/me", middlewareProvider.Auth(), ctrl.GetMeAPI())
	v1.GET("/admin", middlewareProvider.Auth(), middlewareProvider.CheckRoles(shared.RoleAdmin), ctrl.AdminAPI())
	// users := v1.Group("/users")
	// {
	// 	// users.POST("/register", ctrl.RegisterUser())
	// }
}

func (ctrl *HTTPUserController) AdminAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, shared.SimpleResponse("admin API"))
	}
}