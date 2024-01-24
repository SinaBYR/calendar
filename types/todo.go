package types

type Todo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type CreateTodoRequest struct {
	Name   string `json:"name"`
	Done   bool   `json:"done"`
	UserID int64  `json:"userId"`
}

type DeleteTodoRequest struct {
	ID int64 `param:"id"`
}

type CompleteTodoRequest struct {
	ID int64 `param:"id"`
}

type GetTodoRequest struct {
	ID int64 `param:"id"`
}

type GetTodosRequest struct {
	ID int64 `param:"id"`
}
