package main

import (
	"github.com/21amir21/event-booking/db"
	"github.com/21amir21/event-booking/gintemplrenderer"
	"github.com/21amir21/event-booking/routes"
	"github.com/21amir21/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	utils.LoadEnv()

	r := gin.Default()

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning
	r.SetTrustedProxies(nil)

	routes.RegisterRoutes(r)

	r.Run(":8080") // localhost:8080
}
