package main
import (
	"github.com/gin-gonic/gin"
	"github.com/ponsonio/avanticaChallenge/pkg/api/app"
)
var Router *gin.Engine

func main() {
	app.StartApp()
}