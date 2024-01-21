package models

import "database/sql"

type User struct {
	ID        int64          `gorm:"<-:false;primaryKey"`
	UserName  string         `gorm:"type:varchar(100) not null unique"`
	FirstName sql.NullString `gorm:"type:varchar(50)"`
	LastName  sql.NullString `gorm:"type:varchar(50)"`
}

type Todo struct {
	ID     int64  `gorm:"<-:false;primaryKey"`
	Name   string `gorm:"type:varchar(100) not null"`
	Done   bool
	UserID int64 `gorm:"->;<-:create"`
	User   User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
