package main

import (
	"fmt"

	"github.com/halhal23/shoezoo/gateway/router"
)

const (
	defaltPort = ":8080"
)

func main() {
	fmt.Println("hello shoezoo gateway.")
	r := router.NewRouter()
	r.Run(defaltPort)
}
