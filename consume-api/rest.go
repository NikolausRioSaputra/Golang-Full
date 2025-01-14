package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func consumeREST(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendekode data"})
		return
	}

	c.JSON(http.StatusOK, post)
}
