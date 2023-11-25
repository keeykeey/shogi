package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "backend/service"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.GET("/api/koma",func(c *gin.Context) {
    service.GetKoma(c)
  })

  r.GET("/api/komaArrangements/:arrangementId", func(c *gin.Context) {
    id := c.Param("arrangementId")
    service.GetKomaArrangements(c, id)
  })

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}