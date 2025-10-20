package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 19.99},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func deleteAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	idAsNumber, _ := strconv.Atoi(id)

	for _, album := range albums {
		if album.ID == id {
			albums = append(albums[:idAsNumber-1], albums[idAsNumber:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Album eliminado"})
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func postAlbums(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumById)

	router.Run("localhost:8080")
}
