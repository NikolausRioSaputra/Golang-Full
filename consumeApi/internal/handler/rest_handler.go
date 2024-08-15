package handler

import (
	"consume-api/internal/domain"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConsumeRest handles the REST API call
func ConsumeRest(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data"})
		return
	}
	defer resp.Body.Close()


	var post []domain.Post
	body, _ := io.ReadAll(resp.Body)
	if  err = json.Unmarshal(body, &post); err != nil {
		// if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// type Coba struct {
// }

// func (Coba) Read(paw []byte) (nadw int, erradwd error) {
// 	return 0, nil
// }

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
