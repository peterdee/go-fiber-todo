package todos

type CreateTodoRequest struct {
	Text string `json:"text"`
}

type UpdateTodoRequest struct {
	Completed bool `json:"completed"`
	CreateTodoRequest
}
