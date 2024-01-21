package types

type ResponseMessage struct {
	Message string
	Data    any
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}

type UpdateUserRequest struct {
	ID        int64  `param:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type DeleteUserRequest struct {
	ID int64 `param:"id"`
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
