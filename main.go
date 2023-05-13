package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func main() {
	router := gin.Default()

	if err := router.Run(); err != nil {
		panic("Application not Running")

	}

}
