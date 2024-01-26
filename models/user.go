package models

import "database/sql"

type User struct {
	ID        int64          `gorm:"<-:false;primaryKey"`
	UserName  string         `gorm:"type:varchar(100) not null unique"`
	FirstName sql.NullString `gorm:"type:varchar(50)"`
	LastName  sql.NullString `gorm:"type:varchar(50)"`
	Password  string         `gorm:"type:varchar(50) not null"`
}
