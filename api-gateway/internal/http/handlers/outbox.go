package handlers

import (
	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SendMessage godoc
// @Summary Send a new message
// @Description Sends a new message to the specified recipients.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param message body pb.OutboxMessageSentBody true "Message sending request"
// @Success 201 {object} pb.MessageSentRes "Message sent successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Failure 429 {object} string "Too many requests, please try again later"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox [post]
func (h *HTTPHandler) SendMessage(c *gin.Context) {
	req := pb.OutboxMessageSentReq{}
	if err := c.ShouldBindJSON(&req.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if req.SenderId, err = config.GetUserIDByClaims(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// rateLimitKey := fmt.Sprintf("ratelimit:%s", req.SenderId)
	// exists, err := h.RDB.DB.Exists(c.Request.Context(), rateLimitKey).Result()
	if err != nil {
		h.Logger.ERROR.Printf("Error checking rate limit: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check rate limit", "details": err.Error()})
		return
	}

	// if exists == 1 {
	// 	c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, please wait 10 seconds before sending another message"})
	// 	return
	// }

	// err = h.RDB.DB.Set(c.Request.Context(), rateLimitKey, "active", 30*time.Second).Err()
	// if err != nil {
	// 	h.Logger.ERROR.Printf("Error setting rate limit: %v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set rate limit", "details": err.Error()})
	// 	return
	// }

	res, err := h.OS.Send(c.Request.Context(), &req)
	if err != nil {
		h.Logger.ERROR.Printf("Error sending message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetOutboxMessageByID godoc
// @Summary Get outbox message by ID
// @Description Retrieves an outbox message by its ID.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param id path string true "Outbox message ID"
// @Success 200 {object} pb.OutboxMessageGetRes "Outbox message retrieved successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Outbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox/{id} [get]
func (h *HTTPHandler) GetOutboxMessageByID(c *gin.Context) {
	messageId := c.Param("id")

	res, err := h.OS.Get(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error getting outbox message by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get outbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllOutboxMessages godoc
// @Summary Get all outbox messages
// @Description Retrieves all outbox messages for the authenticated user.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param is_archived query bool false "Filter by archived status"
// @Param is_trashed query bool false "Filter by if it is in trash"
// @Param is_draft query bool false "Filter by draft status"
// @Param is_starred query bool false "Filter by starred status"
// @Param sent_from query string false "Filter by sent date (from). syntax: 2024-09-07T12:18:28+00:00"
// @Param sent_to query string false "Filter by sent date (to). syntax: 2024-09-07T12:18:28+00:00"
// @Param page query int false "Page number"
// @Param limit query int false "Number of messages per page"
// @Success 200 {object} pb.OutboxMessagesGetAllRes "Outbox messages retrieved successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox [get]
func (h *HTTPHandler) GetAllOutboxMessages(c *gin.Context) {
	user_id, err := config.GetUserIDByClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	req := &pb.OutboxMessagesGetAllReq{
		SenderId: user_id,
		Body: &pb.OutboxMessagesGetAllBody{
			Query:      c.Query("query"),
			IsArchived: c.Query("is_archived") == "true",
			IsTrashed:  c.Query("is_trashed") == "true",
			IsDraft:    c.Query("is_draft") == "true",
			IsStarred:  c.Query("is_starred") == "true",
			SentFrom:   c.Query("sent_from"),
			SentTo:     c.Query("sent_to"),
		},
		Pagination: &pb.Pagination{
			Skip:  0,
			Limit: 1000,
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

	res, err := h.OS.GetAll(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Printf("Error getting all outbox messages: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get outbox messages", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// MoveOutboxMessageToTrash godoc
// @Summary Move outbox message to trash
// @Description Moves an outbox message to the trash folder or gets it back if it is already in trash.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param id path string true "Outbox message ID"
// @Success 204 {object} string "Outbox message moved to trash successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Outbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox/{id}/trash [put]
func (h *HTTPHandler) MoveOutboxMessageToTrash(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.OS.MoveToTrash(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error moving outbox message to trash: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move outbox message to trash", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Outbox message moved to / retrived back from trash successfully"})
}

// DeleteOutboxMessage godoc
// @Summary Delete outbox message
// @Description Permanently deletes an outbox message.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param id path string true "Outbox message ID"
// @Success 204 {object} string "Outbox message deleted successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Outbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox/{id} [delete]
func (h *HTTPHandler) DeleteOutboxMessage(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.OS.Delete(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error deleting outbox message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete outbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Outbox message deleted successfully"})
}

// StarOutboxMessage godoc
// @Summary Star outbox message
// @Description Stars or unstars an outbox message.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param id path string true "Outbox message ID"
// @Success 200 {object} pb.Void "Outbox message starred/unstarred successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Outbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox/{id}/star [put]
func (h *HTTPHandler) StarOutboxMessage(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.OS.StarMessage(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error starring/unstarring outbox message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to star/unstar outbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Outbox message starred/unstarred successfully"})
}

// ArchiveOutboxMessage godoc
// @Summary Archive outbox message
// @Description Archives or unarchives an outbox message.
// @Tags 06-Outbox
// @Accept json
// @Produce json
// @Param id path string true "Outbox message ID"
// @Success 200 {object} pb.Void "Outbox message archived/unarchived successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Outbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /outbox/{id}/archive [put]
func (h *HTTPHandler) ArchiveOutboxMessage(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.OS.ArchiveMessage(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error archiving/unarchiving outbox message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to archive/unarchive outbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Outbox message archived/unarchived successfully"})
}
