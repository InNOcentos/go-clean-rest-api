package user

import (
	"net/http"

	user "github.com/InNOcentos/go-clean-rest-api/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(*gin.Context)
}

type handler struct {
	useCase user.UseCase
}

func NewHandler(uc user.UseCase) *handler {
	return &handler{
		useCase: uc,
	}
}

func (h *handler) Create(c *gin.Context) {
	var payload user.CreateUserRequest

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.useCase.CreateUser(payload)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}
