package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ponsonio/avanticaChallenge/pkg/api/models"
	"github.com/ponsonio/avanticaChallenge/pkg/api/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var (
	quadrantService services.QuadrantInterface
)

func init() {
	quadrantService = &services.QuadrantService{}
}

func GetQuadrant(context *gin.Context) {

	quadrantID := context.Param("id")
	quadrant, err := (quadrantService).Get(quadrantID)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, *quadrant)

}

func DeleteQuadrant(context *gin.Context) {

	quadrant := models.Quadrant{}
	quadrantID := context.Param("id")

	quadrant.ID, _ = primitive.ObjectIDFromHex(quadrantID)

	err := (quadrantService).Delete(&quadrant)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, "Quadrant deleted")

}

func CreateQuadrant(context *gin.Context) {

	quadrant := models.Quadrant{}

	context.ShouldBindWith(&quadrant, binding.JSON)

	quadrantRes, err := (quadrantService).Create(&quadrant)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, quadrantRes)

}

func UpdateQuadrant(context *gin.Context) {

	quadrant := models.Quadrant{}

	context.ShouldBindWith(&quadrant, binding.JSON)

	quadrantRes, err := (quadrantService).Update(&quadrant)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, quadrantRes)

}
