// backend/internal/auth/user_handler.go
package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"project/ent"
	"project/ent/user"
)

// UserHandler handles user authentication
type UserHandler struct {
	client      *ent.Client
	authService *Service
}

// NewUserHandler creates a new user handler
func NewUserHandler(client *ent.Client, authService *Service) *UserHandler {
	return &UserHandler{
		client:      client,
		authService: authService,
	}
}

// RegisterUser registers a new user
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username already exists
	exists, err := h.client.User.
		Query().
		Where(user.UsernameEQ(input.Username)).
		Exist(c.Request.Context())
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	// Check if email already exists
	exists, err = h.client.User.
		Query().
		Where(user.EmailEQ(input.Email)).
		Exist(c.Request.Context())
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	passwordHash, err := h.authService.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// Create user
	u, err := h.client.User.
		Create().
		SetUsername(input.Username).
		SetPasswordHash(passwordHash).
		SetEmail(input.Email).
		SetRoles([]string{"user"}).
		SetCreatedAt(time.Now()).
		Save(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
		"roles":    u.Roles,
	})
}

// LoginUser authenticates a user and returns a JWT token
func (h *UserHandler) LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by username
	u, err := h.client.User.
		Query().
		Where(user.UsernameEQ(input.Username)).
		Only(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !h.authService.CheckPassword(input.Password, u.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := h.authService.GenerateToken(u.ID, u.Username, u.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       u.ID,
			"username": u.Username,
			"email":    u.Email,
			"roles":    u.Roles,
		},
	})
}