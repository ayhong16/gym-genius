package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/initializers"

	"github.com/gin-gonic/gin"
)

var APIKey string

func init() {
	APIKey = initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/exercises", fetchExercises)
	r.Run()
}

func fetchExercises(c *gin.Context) {
	url := "https://exercisedb.p.rapidapi.com/exercises?limit=0&offset=0"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Add("x-rapidapi-host", "exercisedb.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch exercises"})
		return
	}

	var exercises []Exercise
	if err := json.NewDecoder(resp.Body).Decode(&exercises); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}
	fmt.Printf("Returned %d exercises\n", len(exercises))

	c.JSON(http.StatusOK, exercises)
}
