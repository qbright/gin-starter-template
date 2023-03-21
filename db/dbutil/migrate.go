package dbutil

import (
	"gin-starter-template/modules"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&modules.User{})
}
