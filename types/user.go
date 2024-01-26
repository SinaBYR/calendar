package types

import "database/sql"

type User struct {
	ID        int64          `json:"id"`
	UserName  string         `json:"username"`
	FirstName sql.NullString `json:"firstname"`
	LastName  sql.NullString `json:"lastname"`
}

type UpdateUserRequest struct {
	ID        int64  `param:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type DeleteUserRequest struct {
	ID int64 `param:"id"`
}
