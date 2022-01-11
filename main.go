package main

import (
	"fmt"

	"github.com/antistud/tiptoe_server/routes"
	"github.com/gin-gonic/gin"
)

var err error

func main() {
	if err != nil {
		fmt.Println("status: ", err)
	}
	r := gin.Default()
	routes.SetupRouter(r)
	// running
	r.Run()
}
