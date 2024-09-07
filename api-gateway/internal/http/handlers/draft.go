package handlers

import (
	"net/http"

	"api-gateway/internal/pkg/config"
	pb "api-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateDraft godoc
// @Summary Create a new draft
// @Description Creates a new draft message associated with the authenticated user.
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
	var (
		req pb.DraftCreateUpdateReq
	)

	err := c.ShouldBindJSON(&req.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	claims, err := config.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userId := claims["user_id"].(string)

	req.SenderId = userId

	_, err = h.DS.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &pb.Void{})
}

// UpdateDraft godoc
// @Summary Update an existing draft
// @Description Updates a draft message with the provided ID.
// @Tags 05-Draft
// @Accept json
// @Produce json
// @Param id path string true "Draft ID"
// @Param draft body pb.DraftCreateUpdateReq true "Draft update request"
// @Success 200 {object} pb.Void "Draft updated successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Draft not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /draft/{id} [put]
func (h *HTTPHandler) UpdateDraft(c *gin.Context) {
	var (
		req pb.DraftCreateUpdateReq
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	claims, err := config.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userId := claims["user_id"].(string)

	req.SenderId = userId

	_, err = h.DS.Update(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &pb.Void{})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Draft deleted successfully"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send draft", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
