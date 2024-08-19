package src

import (
	"server/types"

	"github.com/gin-gonic/gin"
)

func StartRouter(db *Database) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, world!"})
	})

	r.GET("/workouts", func(c *gin.Context) {
		workouts, err := db.FetchWorkouts()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch workouts"})
			return
		}
		c.JSON(200, gin.H{"workouts": workouts})
	})

	r.POST("/workout", func(c *gin.Context) {
		var newWorkout types.Workout

		if err := c.ShouldBindJSON(&newWorkout); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		c.JSON(200, gin.H{"message": "Workout created"})
	})

	r.Run()
}
