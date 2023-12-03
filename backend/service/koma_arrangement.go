package service

import (
	"backend/repository"
	"backend/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetKomaArrangements(c *gin.Context, arrangementId string) {
	id, err := strconv.Atoi(arrangementId)
	if err != nil {
		panic("invalid paramater")
	}
	var komaArrangements []repository.KomaArrangements = repository.GetKomaArrangements(uint16(id))

	response.KomaArrangements(c, komaArrangements)
}
