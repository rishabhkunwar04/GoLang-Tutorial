//go mod init goAPI:- giving it the path of the module your code will be in,This command creates a go.mod file in which dependencies you add will be listed for tracking
//go run main.go:- to run project
//go run . :- to run project globally
// go get github.com/gin-gonic/gin  :- downloading gin library

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// client and server communicate via JSON data only
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) //Call Context.IndentedJSON to serialize the struct into JSON and add it to the response
}
func addToAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//adding collected data to list
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)

}
func getAlbumByIdHelp(id string) *album {

	for i, x := range albums {
		if x.ID == id {
			return &albums[i]
		}
	}
	return nil
}
func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	res := getAlbumByIdHelp(id)
	c.IndentedJSON(http.StatusOK, res)

}
func patchAlbums(c *gin.Context) {
	id := c.Param("id")

	for _, x := range albums {
		if x.ID == id {
			x.Title = "Rishabh Kunwar Rajpoot"
			c.IndentedJSON(http.StatusOK, x)
			break

		}
	}
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addToAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.PATCH("/albums/:id", patchAlbums)

	router.Run("localhost:8080")
}
