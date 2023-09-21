package user

import (
	user "github.com/InNOcentos/go-clean-rest-api/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPHandlers(router *gin.Engine, uc user.UseCase) {
	h := NewHandler(uc)

	users := router.Group("/users")
	{
		users.POST("", h.Create)
    users.GET(":id", h.Get)
	}
}
