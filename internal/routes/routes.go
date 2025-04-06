package routes

import (
	"net/http"

	"github.com/Liedsonfsa/API-Challenges-Points-of-Interest/internal/database"
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

	if err = database.RegisterPOI(poi.Name, poi.X, poi.Y); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastro o ponto no banco de dados"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": poi})
}