package handlers

import (
	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"
	rdb "api-gateway/internal/pkg/redis"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// ForgotPassword godoc
// @Summary Forgot passwrod
// @Description Sends a confirmation code to email recovery password
// @Tags 02-Password-recovery
// @Accept json
// @Produce json
// @Param credentials body pb.ByEmail true "User login credentials"
// @Success 200 {object} string ""
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Page not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /forgot-password [POST]
func (h *HTTPHandler) ForgotPassword(c *gin.Context) {
	var req pb.ByEmail
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	user, err := h.US.GetUserByEmail(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		h.Logger.ERROR.Printf("Error getting user by email: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = rdb.SendConfirmationCode(*user.Email, h.RDB, h.Logger)
	if err != nil {
		h.Logger.ERROR.Printf("Error sending confirmation code: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code to email", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmation code sent to your email. Please use your code within 3 minutes."})
}

// RecoverPassword godoc
// @Summary Recover password (Use this one after sending verification code)
// @Description Verifies the code and updates the password
// @Tags 02-Password-recovery
// @Accept json
// @Produce json
// @Param request body pb.UserRecoverPasswordReq true "Recover Password Request"
// @Success 200 {object} string "Password successfully updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Incorrect verification code"
// @Failure 404 {object} string "Verification code expired or email not found"
// @Failure 500 {object} string "Error updating password"
// @Router /recover-password [post]
func (h *HTTPHandler) RecoverPassword(c *gin.Context) {
	var req pb.UserRecoverPasswordReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, code, and new password are required fields."})
		return
	}

	if err := config.IsValidPassword(req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedCode, err := h.RDB.Get(context.Background(), req.Email).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verification code expired or email not found"})
		return
	} else if err != nil {
		h.Logger.ERROR.Printf("Error getting confirmation code from Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email not found in a recovery requests!"})
		return
	}

	if storedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect verification code"})
		return
	}

	hashedPassword, err := config.HashPassword(req.NewPassword)
	if err != nil {
		h.Logger.ERROR.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't hash your password", "details": err.Error()})
		return
	}
	_, err = h.US.ChangeUserPassword(c, &pb.UserRecoverPasswordReq{Email: req.Email, NewPassword: hashedPassword})
	if err != nil {
		h.Logger.ERROR.Printf("Error updating password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password", "details": err.Error()})
		return
	}
	h.RDB.Del(context.Background(), req.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Password successfully updated"})
}

// SendCodeAgain godoc
// @Summary Sends code again if user didn't recieve the code
// @Description Sends a confirmation code to email recovery password again
// @Tags 02-Password-recovery
// @Accept json
// @Produce json
// @Param credentials body pb.ByEmail true "User login credentials"
// @Success 200 {object} string "Code sent"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Page not found"
// @Failure 500 {object} string "Server error"
// @Router /send-code-again [POST]
func (h *HTTPHandler) SendCodeAgain(c *gin.Context) {
	var req pb.ByEmail
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	user, err := h.US.GetUserByEmail(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		h.Logger.ERROR.Printf("Error getting user by email: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = rdb.SendConfirmationCode(*user.Email, h.RDB, h.Logger)
	if err != nil {
		h.Logger.ERROR.Printf("Error sending confirmation code: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code to email", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmation code sent to your email. Please use your code within 3 minutes."})
}
