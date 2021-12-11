package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product "github.com/halhal23/shoezoo/gateway/client/proto/v1"
)

type Handler struct {
	ProductClient product.ProductServiceClient
}

func NewHandler(client product.ProductServiceClient) *Handler {
	return &Handler{
		ProductClient: client,
	}
}

func (h Handler) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello gateway for shoezoo! 2",
		})
	}
}
