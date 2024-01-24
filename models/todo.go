package models

type Todo struct {
	ID     int64  `gorm:"<-:false;primaryKey"`
	Name   string `gorm:"type:varchar(100) not null"`
	Done   bool
	UserID int64 `gorm:"->;<-:create"`
	User   User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
