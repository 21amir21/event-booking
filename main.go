package main

import (
	"github.com/21amir21/event-booking/db"
	"github.com/21amir21/event-booking/routes"
	"github.com/21amir21/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	utils.LoadEnv()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080") // localhost:8080
}
