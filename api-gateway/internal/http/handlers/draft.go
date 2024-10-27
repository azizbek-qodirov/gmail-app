package handlers

import (
	"net/http"

	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateDraft godoc
// @Summary Create a new draft
// @Description Creates a new draft message. Use empty string in order to not to use this field.
// @Tags 05-Draft
// @Accept json
// @Produce json
// @Param draft body pb.DraftCreateUpdateBody true "Draft creation request"
// @Success 201 {object} pb.Void "Draft created successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /draft [post]
func (h *HTTPHandler) CreateDraft(c *gin.Context) {
	var req pb.DraftCreateUpdateReq

	if err := c.ShouldBindJSON(&req.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if req.SenderId, err = config.GetUserIDByClaims(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if _, err = h.DS.Create(c.Request.Context(), &req); err != nil {
		h.Logger.ERROR.Printf("Error creating draft: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Draft successfully created!"})
}

// UpdateDraft godoc
// @Summary Update an existing draft
// @Description Updates a draft message with the provided ID. Use empty string in order to not to use this field.
// @Tags 05-Draft
// @Accept json
// @Produce json
// @Param id path string true "Draft ID"
// @Param draft body pb.DraftCreateUpdateBody true "Draft update request"
// @Success 200 {object} pb.Void "Draft updated successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Draft not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /draft/{id} [put]
func (h *HTTPHandler) UpdateDraft(c *gin.Context) {
	var req pb.DraftCreateUpdateReq
	draft_id := c.Param("id")
	req.SenderId = draft_id

	if err = c.ShouldBindJSON(&req.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if _, err = h.DS.Update(c.Request.Context(), &req); err != nil {
		h.Logger.ERROR.Printf("Error updating draft: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Draft successfully updated!"})
}

// DeleteDraft godoc
// @Summary Delete a draft
// @Description Deletes a draft message with the provided ID.
// @Tags 05-Draft
// @Accept json
// @Produce json
// @Param id path string true "Draft ID"
// @Success 204 {object} string "Draft deleted successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Draft not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /draft/{id} [delete]
func (h *HTTPHandler) DeleteDraft(c *gin.Context) {
	draftId := c.Param("id")

	_, err := h.DS.Delete(c.Request.Context(), &pb.ByID{Id: draftId})
	if err != nil {
		h.Logger.ERROR.Printf("Error deleting draft: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Draft deleted successfully"})
}

// SendDraft godoc
// @Summary Send a draft
// @Description Sends a draft message with the provided ID.
// @Tags 05-Draft
// @Accept json
// @Produce json
// @Param id path string true "Draft ID"
// @Success 200 {object} pb.MessageSentRes "Draft sent successfully"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Draft not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /draft/{id}/send [post]
func (h *HTTPHandler) SendDraft(c *gin.Context) {
	draftId := c.Param("id")

	res, err := h.DS.SendDraft(c.Request.Context(), &pb.ByID{Id: draftId})
	if err != nil {
		h.Logger.ERROR.Printf("Error sending draft: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
