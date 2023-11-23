package response

import (
	"backend/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

func KomaBase(c *gin.Context, komaBases []repository.KomaBase) {
    c.JSON(http.StatusOK, komaBases)
}
