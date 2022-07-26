package controller

import (
	"net/http"
	"strconv"

	"clean-architecture/entity"
	"clean-architecture/usecase"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service usecase.Service
}

// NewUserController creates new user controller
func NewUserController(userService usecase.Service) UserController {
	return UserController{
		service: userService,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := u.service.GetUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, userResponse{user})

}

type userResponse struct {
	User *entity.User `json:"user"`
}
