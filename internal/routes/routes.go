package routes

import (
	"net/http"

	"github.com/Liedsonfsa/API-Challenges-Points-of-Interest/internal/models"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var poi models.POI

	err := context.ShouldBindJSON(&poi)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível obter os dados enviados"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": poi})
}