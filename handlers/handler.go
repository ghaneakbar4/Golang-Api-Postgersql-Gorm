package handlers


import (
	"encoding/json"
	"fmt"
	"gorm/models"
	"net/http"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}
//تزریق وابستگی
func New(db *gorm.DB) handler {
	return handler{db}
}

//متدهای وب سرویس
func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {
	
    var user models.User
	//دریافت اطلاعات از هدر
     json.NewDecoder(r.Body).Decode(&user)
    // افزودن به تیبل 
    if result := h.DB.Create(&user); result.Error != nil {
        fmt.Println(result.Error)
    }

    
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(&user)
}
func (h handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
// لیست کاربران
    if result := h.DB.Find(&users); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}