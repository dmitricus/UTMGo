package controllers

import (
	"github.com/gin-gonic/gin"
	"main/auth/models"
	"net/http"
)

type CreateUserInput struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	IsSuperusers bool   `json:"is_superusers" binding:"required"`
	IsStaff      bool   `json:"is_staff" binding:"required"`
	IsActive     bool   `json:"is_active" binding:"required"`
}

// POST /users
// Создание user
//http://127.0.0.1:8081/auth/create_user/
//{
//	"username":     "root",
//	"password":     "root",
//	"email":        "admin@admin.local",
//	"first_name":    "Бородулин",
//	"last_name":     "Дмитирий",
//	"is_superusers": true,
//	"is_staff":      true,
//	"is_active":     true
//}
func CreateUser(context *gin.Context) {
	var input CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username:     input.Username,
		Password:     input.Password,
		Email:        input.Email,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		IsSuperusers: input.IsSuperusers,
		IsStaff:      input.IsStaff,
		IsActive:     input.IsActive,
	}
	//MainModels.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"users": user})
}
