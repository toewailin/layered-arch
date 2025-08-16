package handlers

import (
	"app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Login handles the user login request
func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the JSON request data to the loginRequest object
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Call service to authenticate the user and get the token
	token, err := h.Service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Return success response with the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register handles the user registration request
func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	// Bind the JSON request data to the registerRequest object
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Call service to register the user
	if err := h.Service.Register(registerRequest.Email, registerRequest.Password, registerRequest.FirstName, registerRequest.LastName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
