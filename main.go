package main

import (
	"gorm/db"
	"gorm/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Product struct {
	ID    uint `gorm:"primaryKey;autoIncrement:true"`
	Code  string
	Price uint
}

func main() {
	//ساخت نمونه
	DB :=db.Init()
	//تزریق به استراکت
	h :=handlers.New(DB)
	//استفاده از روت گوریلا 
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", h.GetAllUsers).Methods(http.MethodGet)
	log.Println("API is running!")
	//راه اندازی سرور
    http.ListenAndServe(":4000", router)

}