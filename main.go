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

	if err := db.DbInit(); err != nil {
		panic("Could not connect to database")
	}

	defer db.Client.Disconnect(db.Ctx)
	defer db.CtxCancel()
	/*
	   List databases
	*/

	r := gin.Default()
	routes.SetupRouter(r)
	// running
	r.Run()
}
