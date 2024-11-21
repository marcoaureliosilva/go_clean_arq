package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

// NewHandler cria um novo handler para ordens
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// ListOrders lista todas as ordens
func (h *Handler) ListOrders(c *gin.Context) {
	var orders []Order
	h.db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
