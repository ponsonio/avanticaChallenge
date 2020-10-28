package app

import (
	"github.com/ponsonio/avanticaChallenge/pkg/api/controllers"
)

func mapUrls() {
	router.GET("/api/maze/health", controllers.Test)

	router.GET("api/maze/spot/:id", controllers.GetSpot)
	router.POST("api/maze/spot/", controllers.CreateSpot)
	router.DELETE("api/maze/spot/:id", controllers.DeleteSpot)
	router.PUT("api/maze/spot/", controllers.UpdateSpot)

	router.GET("api/maze/path/:id", controllers.GetPath)
	router.POST("api/maze/path/", controllers.CreatePath)
	router.DELETE("api/maze/path/:id", controllers.DeletePath)
	router.PUT("api/maze/path/", controllers.UpdatePath)

	router.GET("api/maze/quadrant/:id", controllers.GetQuadrant)
	router.POST("api/maze/quadrant/", controllers.CreateQuadrant)
	router.DELETE("api/maze/quadrant/:id", controllers.DeleteQuadrant)
	router.PUT("api/maze/quadrant/", controllers.UpdateQuadrant)
}
