package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

// CreateAttachment godoc
// @Summary Create a new attachment
// @Description Creates a new attachment associated with an outbox message.
// @Tags 07-Attachments
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Attachment file"
// @Success 201 {object} pb.AttachmentCreateRes "Attachment created successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /attachment [post]
func (h *HTTPHandler) CreateAttachment(c *gin.Context) {
	user_id, err := config.GetUserIDByClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload", "details": err.Error()})
		return
	}
	defer file.Close()

	if header.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size too large. Maximum allowed size is 10MB"})
		return
	}

	uniqueFileName := config.GenRandomNum() + "_" + header.Filename

	uploadInfo, err := h.Minio.Client.PutObject(
		c.Request.Context(),
		h.Minio.DefaultBucket(),
		uniqueFileName,
		file,
		header.Size,
		minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")},
	)
	if err != nil {
		h.Logger.ERROR.Printf("Error uploading file to Minio: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Minio", "details": err.Error()})
		return
	}

	fileUrl := fmt.Sprintf("http://52.77.251.174:9000/%s/%s", h.Minio.DefaultBucket(), uploadInfo.Key)

	req := &pb.AttachmentCreateReq{
		UserId:   user_id,
		FileUrl:  fileUrl,
		FileName: uniqueFileName,
		FileSize: strconv.FormatInt(header.Size, 10),
		MimeType: header.Header.Get("Content-Type"),
	}

	res, err := h.AS.Create(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Printf("Error creating attachment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attachment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetAttachmentsByOutboxID godoc
// @Summary Get attachments for an outbox message
// @Description Retrieves all attachments associated with an outbox message.
// @Tags 07-Attachments
// @Accept json
// @Produce json
// @Param outbox_id path string true "Outbox message ID"
// @Param mime_type query string false "Filter by MIME type"
// @Param created_from query string false "Filter by creation date (from)"
// @Param created_to query string false "Filter by creation date (to)"
// @Param page query int false "Page number"
// @Param limit query int false "Number of attachments per page"
// @Success 200 {object} pb.AttachmentGetAllRes "Attachments retrieved successfully"
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Attachments not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /attachment/{outbox_id} [get]
func (h *HTTPHandler) GetAttachmentsByOutboxID(c *gin.Context) {
	outboxId := c.Param("outbox_id")

	req := &pb.AttachmentGetAllReq{
		OutboxId: outboxId,
	}

	res, err := h.AS.GetAll(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Printf("Error getting attachments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get attachments", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteAttachment godoc
// @Summary Delete an attachment
// @Description Deletes an attachment by its ID.
// @Tags 07-Attachments
// @Accept json
// @Produce json
// @Param id path string true "Attachment ID"
// @Success 204 {object} string "Attachment deleted successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Attachment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /attachment/{id} [delete]
func (h *HTTPHandler) DeleteAttachment(c *gin.Context) {
	attachmentId := c.Param("id")

	res, err := h.AS.Delete(c.Request.Context(), &pb.ByID{Id: attachmentId})
	if err != nil {
		h.Logger.ERROR.Printf("Error deleting attachment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment", "details": err.Error()})
		return
	}
	if err = h.Minio.Client.RemoveObject(c.Request.Context(), h.Minio.DefaultBucket(), res.FileName, minio.RemoveObjectOptions{}); err != nil {
		h.Logger.ERROR.Printf("Error deleting attachment from Minio: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment from Minio", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attachment deleted successfully"})
}

// GetMyUploads godoc
// @Summary Get my uploads
// @Description Retrieves all attachments associated with the authenticated user.
// @Tags 07-Attachments
// @Accept json
// @Produce json
// @Success 200 {object} pb.AttachmentGetAllRes "Attachments retrieved successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Attachments not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /attachment/my-uploads [get]
func (h *HTTPHandler) GetMyUploads(c *gin.Context) {
	user_id, err := config.GetUserIDByClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res, err := h.AS.GetMyUploads(c.Request.Context(), &pb.ByID{Id: user_id})
	if err != nil {
		h.Logger.ERROR.Printf("Error getting my uploads: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get attachments", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
