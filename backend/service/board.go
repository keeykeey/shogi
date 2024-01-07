package service

import (
	"backend/repository"
	"backend/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBoard(c *gin.Context, boardId string) {
	id, err := strconv.Atoi(boardId)
	if err != nil {
		panic("invalid parameter")
	}
	var board repository.BoardBase = repository.GetBoard(byte(id))
	response.BoardBase(c, board)
}