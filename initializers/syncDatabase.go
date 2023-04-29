package initializers

import "jwt-go-system/models"

func SyncDatabase() {
	// 创建表
	DB.AutoMigrate(&models.User{})
}
