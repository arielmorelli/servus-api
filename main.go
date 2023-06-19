package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/arielmorelli/servus-api/api"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Accept any route
	router.Any("/*route", api.DynamicRouting)

	// Deregister
	// router.DELETE("/_remove/:id", getAlbums)
	// router.DELETE("/_remove/:route", getAlbums)
	// router.DELETE("/_remove/:method/:route", getAlbums)

	router.Run("localhost:8080")
}
