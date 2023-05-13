package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusCreated, recipe)
}
func ListRecipeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}
func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Recipe not found",
		})
		return
	}
	recipes[index] = recipe

}
func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)

	if err := router.Run(); err != nil {
		panic("Application not Running")

	}

}
