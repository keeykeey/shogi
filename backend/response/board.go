package response

import (
	"backend/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

func BoardBase(c *gin.Context, boardBases repository.BoardBase) {
    c.JSON(http.StatusOK, boardBases)
}