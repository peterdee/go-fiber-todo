package models

// Todo model structure
type Todo struct {
	Completed bool   `json:"completed"`
	Created   int64  `json:"created"`
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Text      string `json:"text"`
}
