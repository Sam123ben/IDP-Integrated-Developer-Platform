package handlers

import (
	"backend/common/utils"
	"backend/services/fetch_internal_env_details/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalHandler struct {
	Repo *repository.InternalRepository
}

func NewInternalHandler(repo *repository.InternalRepository) *InternalHandler {
	return &InternalHandler{Repo: repo}
}

// GetInternalEnvDetails godoc
// @Summary Get internal environment details
// @Description Retrieves all internal environment details for Product 1
// @Tags internal-env-details
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/internal-env-details [get]
func (h *InternalHandler) GetInternalEnvDetails(c *gin.Context) {
	products, err := h.Repo.GetInternalEnvDetails()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve internal environment details")
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
