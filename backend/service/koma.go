package service

import (
	"backend/repository"
	"backend/response"
	"github.com/gin-gonic/gin"
)

func GetKoma(c *gin.Context) {
	var komaBases []repository.KomaBase = repository.GetKomas()
	response.KomaBase(c, komaBases)
}