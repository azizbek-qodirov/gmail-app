package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags 03-User-profile
// @Accept json
// @Produce json
// @Success 200 {object} pb.UserGetRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile [GET]
func (h *HTTPHandler) Profile(c *gin.Context) {
	claims, err := config.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	email := claims["email"].(string)

	user, err := h.US.GetUserByEmail(c, &pb.ByEmail{Email: email})
	if err != nil {
		h.Logger.ERROR.Printf("Error getting user by email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Server error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary Update user's profile. In order not to update any field, leave it as it is like "string" or ""
// @Description Update user's data
// @Tags 03-User-profile
// @Accept json
// @Produce json
// @Param request body pb.UserUpdateBody true "User Request Body"
// @Success 200 {object} string "User updated successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Error updating User"
// @Security BearerAuth
// @Router /user [PUT]
func (h *HTTPHandler) UpdateProfile(c *gin.Context) {
	var req pb.UserUpdateReq
	if err := c.BindJSON(&req.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid request payload", "details": err.Error()})
		return
	}
	claims, err := config.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	id := claims["user_id"].(string)

	req.Id = id
	_, err = h.US.UpdateUser(c, &req)
	if err != nil {
		h.Logger.ERROR.Printf("Error updating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// SetProfilePicture godoc
// @Summary Set a profile picture
// @Description Adds a profile image to user.
// @Tags 03-User-profile
// @Accept multipart/mixed
// @Produce json
// @Param image formData file false "Profile image"
// @Success 200 {object} string "Profile image is added"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /set-pfp [POST]
func (h *HTTPHandler) SetPFP(c *gin.Context) {
	wd, err := os.Getwd()
	if err != nil {
		h.Logger.ERROR.Printf("Error getting working directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get working directory", "details": err.Error()})
		return
	}

	tempFolderPath := filepath.Join(wd, "temp/profile-pictures/")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err = os.MkdirAll(tempFolderPath, 0755); err != nil {
		h.Logger.ERROR.Printf("Error creating directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create directory", "details": err.Error()})
		return
	}

	filepath := fmt.Sprintf("%s%s", tempFolderPath, file.Filename)
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		h.Logger.ERROR.Printf("Error saving file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't save file", "details": err.Error()})
		return
	}

	claims, err := config.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	email := claims["email"].(string)

	info, err := h.Minio.Client.FPutObject(
		c,
		h.Minio.DefaultBucket(),
		fmt.Sprintf("pfp-%s", strings.Split(email, "@")[0]),
		filepath,
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)
	if err != nil {
		h.Logger.ERROR.Printf("Error putting object to Minio: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't put object to minio", "details": err.Error()})
		return
	}
	imgurl := fmt.Sprintf("http://52.77.251.174:9000/%s/%s", h.Minio.DefaultBucket(), info.Key)

	_, err = h.US.ChangeUserPFP(c, &pb.UserChangePFPReq{Email: email, PhotoUrl: imgurl})
	if err != nil {
		h.Logger.ERROR.Printf("Error updating profile picture in database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update pfp in database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile image is added"})
	go func() {
		time.Sleep(30 * time.Second)
		os.Remove(filepath)
	}()
}

// ChangePassword godoc
// @Summary Change password
// @Description Updates the password to new one
// @Tags 03-User-profile
// @Accept json
// @Produce json
// @Param request body pb.UserChangePasswordReq true "Change Password Request"
// @Success 200 {object} string "Password successfully updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Incorrect verification code"
// @Failure 404 {object} string "Verification code expired or email not found"
// @Failure 500 {object} string "Error updating password"
// @Security BearerAuth
// @Router /user-password [PUT]
func (h *HTTPHandler) UpdatePassword(c *gin.Context) {
	var req pb.UserChangePasswordReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if req.Email == "" || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and new password are required fields."})
		return
	}

	if err := config.IsValidPassword(req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{"message": "Password successfully updated"})
}

func (h *HTTPHandler) GetByID(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	user, err := h.US.GetUserByID(c, id)
	if err != nil {
		h.Logger.ERROR.Printf("Error getting user by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't get the user": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteAccount godoc
// @Summary Delete user account
// @Description Deletes the user's account permanently.
// @Tags 03-User-profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "Account deleted successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Router /delete-account [delete]
func (h *HTTPHandler) DeleteAccount(c *gin.Context) {
	claims, err := config.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	refreshToken, err := config.GetRefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	userId := claims["user_id"].(string)

	_, err = h.US.DeleteUser(c, &pb.ByID{Id: userId})
	if err != nil {
		h.Logger.ERROR.Printf("Error deleting user account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete account", "details": err.Error()})
		return
	}

	err = h.RDB.Set(context.Background(), refreshToken, "", time.Hour*24*7).Err()
	if err != nil {
		h.Logger.ERROR.Printf("Error blacklisting refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}