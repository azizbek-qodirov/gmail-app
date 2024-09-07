package handlers

import (
	"net/http"
	"strconv"

	pb "api-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateAttachment godoc
// @Summary Create a new attachment
// @Description Creates a new attachment associated with an outbox message.
// @Tags 07-Attachments
// @Accept multipart/form-data
// @Produce json
// @Param outbox_id formData string true "Outbox message ID"
// @Param file formData file true "Attachment file"
// @Success 201 {object} pb.AttachmentCreateRes "Attachment created successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /attachment [post]
func (h *HTTPHandler) CreateAttachment(c *gin.Context) {
	outboxId := c.PostForm("outbox_id")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload", "details": err.Error()})
		return
	}
	defer file.Close()

	// File size validation (adjust as needed)
	if header.Size > 5*1024*1024 { // 5MB limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size too large. Maximum allowed size is 5MB"})
		return
	}

	// ... (Your file upload logic to storage, e.g., Minio, S3) ...

	// Assuming you have uploaded the file and have the file URL
	fileUrl := "http://your-storage-service/path/to/file" // Replace with actual URL

	req := &pb.AttachmentCreateReq{
		OutboxId: outboxId,
		FileUrl:  fileUrl,
		FileName: header.Filename,
		FileSize: strconv.FormatInt(header.Size, 10),
		MimeType: header.Header.Get("Content-Type"),
	}

	res, err := h.AS.Create(c.Request.Context(), req)
	if err != nil {
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
		OutboxId:    outboxId,
		MimeType:    c.Query("mime_type"),
		CreatedFrom: c.Query("created_from"),
		CreatedTo:   c.Query("created_to"),
		Pagination: &pb.Pagination{
			Skip:  0,
			Limit: 10,
		},
	}

	if c.Query("page") != "" {
		page, err := strconv.ParseInt(c.Query("page"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
		req.Pagination.Skip = (page - 1) * req.Pagination.Limit
	}

	if c.Query("limit") != "" {
		limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		req.Pagination.Limit = limit
	}

	res, err := h.AS.GetAll(c.Request.Context(), req)
	if err != nil {
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

	_, err := h.AS.Delete(c.Request.Context(), &pb.ByID{Id: attachmentId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Attachment deleted successfully"})
}
