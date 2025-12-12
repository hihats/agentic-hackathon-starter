package handlers

import (
	"net/http"

	"app/models"
	"app/services"

	"github.com/gin-gonic/gin"
)

// ShuffleHandler handles shuffle-related HTTP requests
type ShuffleHandler struct {
	service *services.ShuffleService
}

// NewShuffleHandler creates a new shuffle handler
func NewShuffleHandler() *ShuffleHandler {
	return &ShuffleHandler{
		service: services.NewShuffleService(),
	}
}

// Shuffle handles POST /api/shuffle
func (h *ShuffleHandler) Shuffle(c *gin.Context) {
	var req models.ShuffleRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	// Validate participants count (minimum 4)
	if len(req.Participants) < 4 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: "参加者は4名以上必要です",
		})
		return
	}

	// Validate that at least one of group_size or num_groups is provided
	if req.GroupSize == nil && req.NumGroups == nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: "グループサイズまたはグループ数を指定してください",
		})
		return
	}

	// Validate group_size if provided
	if req.GroupSize != nil && *req.GroupSize < 1 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: "グループサイズは1以上を指定してください",
		})
		return
	}

	// Validate num_groups if provided
	if req.NumGroups != nil && *req.NumGroups < 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: "グループ数は2以上を指定してください",
		})
		return
	}

	// Call shuffle service
	groups, err := h.service.Shuffle(req.Participants, req.GroupSize, req.NumGroups)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "shuffle_error",
			Message: err.Error(),
		})
		return
	}

	// Validate minimum 2 groups
	if len(groups) < 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: "2グループ以上に分ける必要があります",
		})
		return
	}

	c.JSON(http.StatusOK, models.ShuffleResponse{
		Groups: groups,
	})
}
