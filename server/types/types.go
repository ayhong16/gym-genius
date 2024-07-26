package types

type Exercise struct {
	BodyPart         string   `json:"bodyPart"`
	Equipment        string   `json:"equipment"`
	GifURL           string   `json:"gifUrl"`
	ID               string   `json:"id"`
	Instructions     []string `json:"instructions"`
	Name             string   `json:"name"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Target           string   `json:"target"`
}
