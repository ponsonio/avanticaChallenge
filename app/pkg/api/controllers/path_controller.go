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
	pathService services.PathServiceInterface
)

func init() {
	pathService = &services.PathService{}
}

func GetPath(context *gin.Context) {

	pathID := context.Param("id")
	path, err := (pathService).Get(pathID)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, *path)

}

func DeletePath(context *gin.Context) {

	path := models.Path{}
	pathID := context.Param("id")

	path.ID, _ = primitive.ObjectIDFromHex(pathID)

	err := (pathService).Delete(&path)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, "path deleted")

}

func CreatePath(context *gin.Context) {

	path := models.Path{}

	context.ShouldBindWith(&path, binding.JSON)

	pathRes, err := (pathService).Create(&path)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, pathRes)

}

func UpdatePath(context *gin.Context) {

	path := models.Path{}

	context.ShouldBindWith(&path, binding.JSON)

	pathRes, err := (pathService).Update(&path)

	if err != nil {
		SetResponseWithError(context, err)
		return
	}

	SetResponseWithBody(context, http.StatusOK, pathRes)

}
