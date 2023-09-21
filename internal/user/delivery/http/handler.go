package user

import (
	"net/http"

	user "github.com/InNOcentos/go-clean-rest-api/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(*gin.Context)
	Get(*gin.Context)
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
  ctx := c.Request.Context()

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.useCase.CreateUser(ctx, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *handler) Get(c *gin.Context) {
  ctx := c.Request.Context()

  id := c.Param("id")

	user, err := h.useCase.GetUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
