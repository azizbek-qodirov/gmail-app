package handlers

import (
	"net/http"

	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"

	"strconv"

	"github.com/gin-gonic/gin"
)

// GetInboxMessageByID godoc
// @Summary Get inbox message by ID
// @Description Retrieves an inbox message by its ID.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 200 {object} pb.InboxMessageGetRes "Inbox message retrieved successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id} [get]
func (h *HTTPHandler) GetInboxMessageByID(c *gin.Context) {
	messageId := c.Param("id")

	res, err := h.IS.GetByID(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error getting inbox message by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllInboxMessages godoc
// @Summary Get all inbox messages
// @Description Retrieves all inbox messages for the authenticated user.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param sender_id query string false "Filter by sender ID"
// @Param type query string false "Filter by message type (to, cc, bcc)"
// @Param is_spam query bool false "Filter by spam status"
// @Param is_archived query bool false "Filter by archived status"
// @Param is_starred query bool false "Filter by starred status"
// @Param is_trashed query bool false "Filter by if it is in trash"
// @Param sent_from query string false "Filter by sent date (from)"
// @Param sent_to query string false "Filter by sent date (to)"
// @Param unread_only query bool false "Filter by unread messages only"
// @Param page query int false "Page number"
// @Param limit query int false "Number of messages per page"
// @Success 200 {object} pb.InboxMessagesGetAllRes "Inbox messages retrieved successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox [get]
func (h *HTTPHandler) GetAllInboxMessages(c *gin.Context) {
	userId, err := config.GetUserIDByClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	req := &pb.InboxMessageGetAllReq{
		ReceiverId: userId,
		Body: &pb.InboxMessageGetAllBody{
			Query:      c.Query("query"),
			SenderId:   c.Query("sender_id"),
			Type:       c.Query("type"),
			IsSpam:     c.Query("is_spam") == "true",
			IsArchived: c.Query("is_archived") == "true",
			IsStarred:  c.Query("is_starred") == "true",
			IsTrashed:  c.Query("is_trashed") == "true",
			SentFrom:   c.Query("sent_from"),
			SentTo:     c.Query("sent_to"),
			UnreadOnly: c.Query("unread_only") == "true",
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

	res, err := h.IS.GetAll(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Printf("Error getting all inbox messages: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inbox messages", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// MoveInboxMessageToTrash godoc
// @Summary Move inbox message to trash
// @Description Moves an inbox message to the trash folder.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 204 {object} string "Inbox message moved to trash successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id}/trash [put]
func (h *HTTPHandler) MoveInboxMessageToTrash(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.IS.MoveToTrash(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error moving inbox message to trash: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move inbox message to trash", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inbox message moved to trash successfully"})
}

// DeleteInboxMessage godoc
// @Summary Delete inbox message
// @Description Permanently deletes an inbox message.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 204 {object} string "Inbox message deleted successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id} [delete]
func (h *HTTPHandler) DeleteInboxMessage(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.IS.Delete(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error deleting inbox message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inbox message deleted successfully"})
}

// MarkInboxMessageAsRead godoc
// @Summary Mark inbox message as read
// @Description Marks an inbox message as read.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 200 {object} pb.Void "Inbox message marked as read successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id}/read [put]
func (h *HTTPHandler) MarkInboxMessageAsRead(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.IS.MarkAsRead(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error marking inbox message as read: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark inbox message as read", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inbox message marked as read successfully"})
}

// MarkInboxMessageAsSpam godoc
// @Summary Mark inbox message as spam/unspam
// @Description Marks an inbox message as spam or unspam if alreadys spammed.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 200 {object} pb.Void "Inbox message marked as spam successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id}/spam [put]
func (h *HTTPHandler) MarkInboxMessageAsSpam(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.IS.MarkAsSpam(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error marking inbox message as spam: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark inbox message as spam", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inbox message marked as spam successfully"})
}

// StarInboxMessage godoc
// @Summary Star inbox message
// @Description Stars or unstars an inbox message.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 200 {object} pb.Void "Inbox message starred/unstarred successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id}/star [put]
func (h *HTTPHandler) StarInboxMessage(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.IS.StarMessage(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error starring/unstarring inbox message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to star/unstar inbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inbox message starred/unstarred successfully"})
}

// ArchiveInboxMessage godoc
// @Summary Archive inbox message
// @Description Archives or unarchives an inbox message.
// @Tags 04-Inbox
// @Accept json
// @Produce json
// @Param id path string true "Inbox message ID"
// @Success 200 {object} pb.Void "Inbox message archived/unarchived successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Inbox message not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /inbox/{id}/archive [put]
func (h *HTTPHandler) ArchiveInboxMessage(c *gin.Context) {
	messageId := c.Param("id")

	_, err := h.IS.ArchiveMessage(c.Request.Context(), &pb.ByID{Id: messageId})
	if err != nil {
		h.Logger.ERROR.Printf("Error archiving/unarchiving inbox message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to archive/unarchive inbox message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inbox message archived/unarchived successfully"})
}
