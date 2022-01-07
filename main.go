package main

import (
	"img-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", routes.Home)
	r.GET("/images/:name", routes.Images)
	r.GET("/images", routes.Img)

	r.Run()

}
