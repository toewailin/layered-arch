package handlers

import (
	"app/models"
	"app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := h.Service.UpdateUser(uint(id), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}