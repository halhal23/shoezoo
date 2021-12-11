package router

import (
	"log"

	"github.com/gin-gonic/gin"
	product "github.com/halhal23/shoezoo/gateway/client/proto/v1"
	"github.com/halhal23/shoezoo/gateway/handler"
	"github.com/halhal23/shoezoo/gateway/middleware"
	"google.golang.org/grpc"
)

const (
	Ping        = "/ping"
	ShowProduct = "/products/:id"
)

func NewRouter() *gin.Engine {
	e := gin.Default()
	e.Use(middleware.Cors())

	// client for product service
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faild to connect product service: %v\n", err)
	}
	client := product.NewProductServiceClient(conn)

	h := handler.NewHandler(client)

	// router
	e.GET(Ping, h.Ping())
	e.GET(ShowProduct, h.ShowProduct())

	return e
}
