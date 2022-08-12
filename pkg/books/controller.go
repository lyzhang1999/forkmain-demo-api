package books

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-gin-api-medium/pkg/common/models"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)

	addFirstBook(db)
}

func addFirstBook(db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	var book = models.Book{}
	var count int64 = 0
	h.DB.Model(&models.Book{}).Where("id = ?", 1).Count(&count)

	if count == 0 {
		book.Title = "War and Peace"
		book.Author = "Leo Tolstoy"
		book.Description = "War and Peace is a literary work mixed with chapters on history and philosophy by the Russian author Leo Tolstoy. It was first published serially, then published in its entirety in 1869."
		h.DB.Create(&book)
	}
}
