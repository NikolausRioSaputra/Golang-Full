package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Post struct {
// 	UserID int    `json:"userId"`
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// }

// func consumeRest(c *gin.Context) {
// 	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "gak dapat data woi"})
// 		return
// 	}
// 	//Baris ini memastikan bahwa body respons akan ditutup setelah selesai diproses untuk membebaskan sumber daya.
// 	defer resp.Body.Close()

// 	var post Post
// 	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "gk dapat data saya"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, post)
// }
