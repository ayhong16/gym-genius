package types

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	BodyPart         string    `json:"bodyPart"`
	Equipment        string    `json:"equipment"`
	GifURL           string    `json:"gifUrl"`
	ID               uuid.UUID `json:"id"`
	Instructions     []string  `json:"instructions"`
	Name             string    `json:"name"`
	SecondaryMuscles []string  `json:"secondaryMuscles"`
	Target           string    `json:"target"`
}

type Workout struct {
	ID                uuid.UUID  `json:"id"`
	Name              string     `json:"name"`
	NumberOfExercises int        `json:"numberOfExercises"`
	EstimatedDuration int        `json:"estimatedDuration"`
	Exercises         []Exercise `json:"exercises"`
	CreationDate      time.Time  `json:"creationDate"`
}
