package db

import (
	"gorm/models"
	"log"
//درایور پستگرس
	"gorm.io/driver/postgres"
	//پکیج gorm
	"gorm.io/gorm"
)

//ساخت یک نمونه از orm 
func Init() *gorm.DB {
	//رشته اتصال به دیتابیس پستگرس
	dbURL := "postgres://postgres:aaaa@1234@localhost:5432/gorm"
   //باز کردن ارتباط با دیتابیس
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
    // ایجاد شدن تیبل ها در دیتابیس
	db.AutoMigrate(&models.User{})

	return db
}