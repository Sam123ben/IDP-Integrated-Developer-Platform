package handlers

import (
	"backend/services/fetch_infra_types/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfraHandler struct {
	Repo *repository.InfraRepository
}

func NewInfraHandler(repo *repository.InfraRepository) *InfraHandler {
	return &InfraHandler{Repo: repo}
}

// GetInfraTypes godoc
// @Summary Get infrastructure types
// @Description Retrieves all infrastructure types (e.g., INTERNAL, CUSTOMER)
// @Tags Infra Types
// @Produce json
// @Success 200 {array} models.InfraType
// @Failure 500 {object} map[string]string "error"
// @Router /api/infra-types [get]
func (h *InfraHandler) GetInfraTypes(c *gin.Context) {
	infraTypes, err := h.Repo.GetAllInfraTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve infrastructure types"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"infraTypes": infraTypes})
}
