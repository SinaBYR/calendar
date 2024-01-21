package db

import "example.com/golang-crud-gorm/models"

func Migrate() error {
	return DB().Migrator().AlterColumn(&models.Todo{}, "UserID")
}
