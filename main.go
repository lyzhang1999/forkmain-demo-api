package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-gin-api-medium/pkg/books"
	"github.com/hellokvn/go-gin-api-medium/pkg/common/db"
)

func main() {

	port := os.Getenv("Port")
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		os.Getenv("DatabaseUser"),
		os.Getenv("DatabasePassword"),
		os.Getenv("DatabaseHost"),
		os.Getenv("DatabaseName"),
	)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/env", func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("ENV_NAME"))
	})

	r.Run("0.0.0.0:" + port)
}
