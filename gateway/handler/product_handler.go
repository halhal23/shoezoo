package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	product "github.com/halhal23/shoezoo/gateway/client/proto/v1"
)

func (h *Handler) ShowProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Error 001: %v", err))
			return
		}
		res, err := h.ProductClient.ShowProduct(c, &product.ShowProductRequest{Id: int32(id)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error 002: %v", err))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":      res.Product.Id,
			"message": "successfully",
		})
	}
}
