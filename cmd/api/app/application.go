package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pos_system/internal/config"
	"pos_system/internal/database"
)

func New() error {
	dbConfig := config.ReadDBConfig()

	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		fmt.Println("Error conectando a la base de datos:", err)
		return err
	}

	defer db.Close()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	return router.Run("localhost:8080")
}
