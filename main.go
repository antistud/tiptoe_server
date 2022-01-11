package main

import (
	"fmt"

	"github.com/antistud/tiptoe_server/db"
	"github.com/antistud/tiptoe_server/routes"
	"github.com/gin-gonic/gin"
)

var err error

func main() {
	if err != nil {
		fmt.Println("status: ", err)
	}

	ctx, cancel, err := db.DbInit()
	if err != nil {
		panic("Could not connect to database")
	}

	defer db.Client.Disconnect(ctx)
	defer cancel()
	/*
	   List databases
	*/

	r := gin.Default()
	routes.SetupRouter(r)
	// running
	r.Run()
}
