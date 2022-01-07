package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//album represents data about a record album
type album struct {
	ID		string 	`json:"id"`
	Title	string 	`json:"title"`
	Artist	string 	`json:"artist"`
	Price	float64	`json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:6000")
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album
	//Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID locates the album with the matching id and returns
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//loop over list of albums
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

/* Testing GET: 
curl http://localhost:6000/albums
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
*/

/* Testing POST:
curl http://localhost:6000/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
*/

/* Testing GET by ID: 
curl http://localhost:6000/albums/2
*/