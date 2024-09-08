package handlers

import (
	"api-gateway/internal/http/token"
	"api-gateway/internal/pkg/config"
	rdb "api-gateway/internal/pkg/redis"
	"time"

	"context"
	"net/http"

	pb "api-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, username, and password
// @Tags 01-Authentication
// @Accept json
// @Produce json
// @Param user body pb.UserCreateReqForSwagger true "User registration request"
// @Success 201 {object} string "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /register [post]
func (h *HTTPHandler) Register(c *gin.Context) {
	var req pb.UserCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	isEmailExists, err := h.US.IsEmailExists(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		h.Logger.ERROR.Printf("Error checking email existence: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't check email existance", "details": err.Error()})
		return
	}
	if isEmailExists.Exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	if err := config.IsValidPassword(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := config.HashPassword(req.Password)
	if err != nil {
		h.Logger.ERROR.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
		return
	}

	req.Password = hashedPassword

	err = rdb.SendConfirmationCode(req.Email, h.RDB, h.Logger)
	if err != nil {
		h.Logger.ERROR.Printf("Error sending confirmation code: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code", "err": err.Error()})
		return
	}

	_, err = h.US.CreateUser(c, &req)
	if err != nil {
		h.Logger.ERROR.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
		return
	}

	h.Logger.INFO.Println("New account registered to the system: ", req.Email)
	c.JSON(http.StatusOK, gin.H{"message": "Your account has been registered. Please check your email for a confirmation link. You have 3 minutes to confirm your account."})
}

// ConfirmRegistration godoc
// @Summary Confirm registration with code
// @Description Confirms a user's registration using the code sent to their email.
// @Tags 01-Authentication
// @Accept json
// @Produce json
// @Param confirmation body pb.ConfirmUserReq true "Confirmation request"
// @Success 200 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Incorrect verification code"
// @Failure 404 {object} string "Verification code expired or email not found"
// @Router /confirm-registration [post]
func (h *HTTPHandler) ConfirmRegistration(c *gin.Context) {
	var req pb.ConfirmUserReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	storedCode, err := h.RDB.Get(context.Background(), req.Email).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verification code expired or email not found"})
		return
	} else if err != nil {
		h.Logger.ERROR.Printf("Error getting confirmation code from Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email not found in confirmation requests!"})
		return
	}

	if storedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect verification code"})
		return
	}

	_, err = h.US.ConfirmUser(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		h.Logger.ERROR.Printf("Error confirming user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error confirming user", "details": err.Error()})
		return
	}

	user, err := h.US.GetUserSecurityByEmail(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		h.Logger.ERROR.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user", "details": err.Error()})
		return
	}

	tokens := token.GenerateJWTToken(user.Id, user.Email)

	h.RDB.Del(context.Background(), req.Email)

	c.JSON(http.StatusOK, tokens)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with email and password
// @Tags 01-Authentication
// @Accept json
// @Produce json
// @Param credentials body pb.LoginReq true "User login credentials"
// @Success 200 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid email or password"
// @Router /login [post]
func (h *HTTPHandler) Login(c *gin.Context) {
	req := pb.LoginReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	user, err := h.US.GetUserSecurityByEmail(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		h.Logger.ERROR.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User registered with this email not found"})
		return
	}

	if !config.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !user.IsConfirmed {
		err = rdb.SendConfirmationCode(req.Email, h.RDB, h.Logger)
		if err != nil {
			h.Logger.ERROR.Printf("Error sending confirmation code: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code", "err": err.Error()})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Your account is not verified. Please check your email for a confirmation link."})
		return
	}

	tokens := token.GenerateJWTToken(user.Id, user.Email)

	c.JSON(http.StatusOK, tokens)
}

// Logout godoc
// @Summary Logout a user
// @Description Logs out the user by blacklisting their refresh token.
// @Tags 01-Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "Successfully logged out"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Router /logout [post]
func (h *HTTPHandler) Logout(c *gin.Context) {
	refreshToken, err := config.GetRefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = h.RDB.Set(context.Background(), refreshToken, "", time.Hour*24*7).Err()
	if err != nil {
		h.Logger.ERROR.Printf("Error blacklisting refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
