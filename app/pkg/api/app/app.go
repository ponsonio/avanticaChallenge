package app

import (
	"github.com/Kamva/mgm/v3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	router = gin.Default()
)

func init() {
	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, "maze", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed initializing the database %v", err)
		panic(err)
	}
	log.Println("conected to db...")

	//Here I should instatiate the controllers with the db (mgm), so they are easy tets using mocks
}

func StartApp() {
	mapUrls()
	//this should be an env vas
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed initializing the router %v", err)
		panic(err)
	}



}
