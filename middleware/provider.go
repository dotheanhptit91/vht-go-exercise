package middleware

import "github.com/gin-gonic/gin"

type IMiddlewareProvider interface {
	Auth() gin.HandlerFunc
	CheckRoles(roles ...string) gin.HandlerFunc
}

type MiddlewareProvider struct {
	tokenInstropecter ITokenIntrospector
}

func NewMiddlewareProvider(tokenInstropecter ITokenIntrospector) *MiddlewareProvider {
	return &MiddlewareProvider{
		tokenInstropecter: tokenInstropecter,
	}
}

func (p *MiddlewareProvider) Auth() gin.HandlerFunc {
	return Auth(p.tokenInstropecter)
}

func (p *MiddlewareProvider) CheckRoles(roles ...string) gin.HandlerFunc {
	return CheckRoles(roles...)
}