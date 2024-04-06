package testmain

/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type album struct {
	ID int `json:"id"`
}

var albums = []album{
	{ID: 1},
	{ID: 2},
}

func HandlerAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func HandlerCreateAlbum(ctx *gin.Context) {
	var newAlbum album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func HandlerGetAlbumID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	for _, albumSelected := range albums {
		if albumSelected.ID == idInt {
			ctx.IndentedJSON(http.StatusOK, albumSelected)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/albums", HandlerAlbums)
	router.GET("/albums/:id", HandlerGetAlbumID)
	router.POST("/create", HandlerCreateAlbum)

	router.Run("localhost:8080")
}
*/
