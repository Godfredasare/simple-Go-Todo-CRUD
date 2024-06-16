package dto

type TodoDTO struct {
	Task        string `json:"task,omitempty" bson:"task,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}
