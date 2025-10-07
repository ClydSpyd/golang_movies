package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ClydSpyd/golang_movies/Server/MagicStreamServer/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/ping/*name", func(c *gin.Context) {
		name := c.Param("name")
		response := gin.H{
			"message": "pong",
			"bing":    "bong",
		}
		if name != "/" {
			response["name"] = name
		}
		c.JSON(http.StatusOK, response)
	})

	r.GET("/movies", controllers.GetMovies())
	r.GET("/movies/:movie_id", controllers.GetMovieById())

	if err := r.Run(":0808"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
