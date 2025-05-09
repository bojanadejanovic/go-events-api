package routes

import (
	"fmt"
	"net/http"

	"bojana.dev/api/models"
	"bojana.dev/api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		utils.HandleValidationError(err, context)
		return
	}

	existingUser, err := models.GetUserByEmail(user.Email)
	if err != nil || existingUser.ID != 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User with this email already exists"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"messsage": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		utils.HandleValidationError(err, context)
		return
	}

	err := user.ValidateCredentials()
	if err != nil {
		fmt.Println("error validating credentials: ", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to authenticate user"})
		return
	}

	login, err := models.SaveLogin(user.ID, token)
	if err != nil {
		fmt.Println("Error saving login: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save login"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "login_id": login.ID})
}
