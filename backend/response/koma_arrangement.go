package response

import (
	"backend/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

func KomaArrangements(c *gin.Context, arrangements []repository.KomaArrangements) {
	c.JSON(http.StatusOK, arrangements)
}
