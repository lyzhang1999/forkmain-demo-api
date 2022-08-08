package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-gin-api-medium/pkg/books"
	"github.com/hellokvn/go-gin-api-medium/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {

	viper.AutomaticEnv()

	port := viper.GetString("Port")
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("DatabaseUser"),
		viper.GetString("DatabasePassword"),
		viper.GetString("DatabaseHost"),
		viper.GetString("DatabasePort"),
		viper.GetString("DatabaseName"),
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
