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
	spotService services.SpotServiceInterface
)

func init() {
	spotService = &services.SpotService{}
}

func GetSpot(context *gin.Context) {

	spotID := context.Param("id")
	spot, err := (spotService).Get(spotID)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, *spot)

}

func DeleteSpot(context *gin.Context) {

	spot := models.Spot{}
	spotID := context.Param("id")

	spot.ID, _ = primitive.ObjectIDFromHex(spotID)

	err := (spotService).Delete(&spot)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, "spot deleted")

}

func CreateSpot(context *gin.Context) {

	spot := models.Spot{}

	context.ShouldBindWith(&spot, binding.JSON)

	spotRes, err := (spotService).Create(&spot)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, spotRes)

}

func UpdateSpot(context *gin.Context) {

	spot := models.Spot{}

	context.ShouldBindWith(&spot, binding.JSON)

	spotRes, err := (spotService).Update(&spot)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, spotRes)

}
