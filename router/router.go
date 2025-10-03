package router

import (
	"net/http"

	db "example.com/go-test-api/db"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", postMovie)
    r.PUT("/movies/:id", putMovie)
	r.DELETE("/movies/:id", deleteMovie)

	return r
}

// -------------------- Handlers --------------------

func postMovie(ctx *gin.Context) {
	var movie db.Movie
	if err := ctx.Bind(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := db.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"movie": res})
}

func getMovies(ctx *gin.Context) {
	res, err := db.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"movies": res})
}

func getMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"movie": res})
}

func putMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	// Ambil data movie lama
	existingMovie, err := db.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Bind JSON baru
	var updatedData db.Movie
	if err := ctx.Bind(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field
	existingMovie.Name = updatedData.Name
	existingMovie.Description = updatedData.Description

	// Simpan ke DB
	res, err := db.UpdateMovie(existingMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"movie": res})
}

func deleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := db.DeleteMovie(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}
