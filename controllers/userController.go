package controllers

import (
	"mi-proyecto/config"
	"mi-proyecto/models"
	"mi-proyecto/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.RespondWithError(c, http.StatusNotFound, "User not found")
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, "Error search user")
		}
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Save(&user).Error; err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error update user")
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Delete(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.RespondWithError(c, http.StatusNotFound, "User not found")
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, "Error to delete user")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func LoginUser(c *gin.Context) {
	var loginInput struct {
		Email    string `json:"email,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password" validate:"required"`
	}
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	var Field string

	if loginInput.Username != "" {
		Field = "Name = ?"
	} else if loginInput.Email != "" {
		Field = "Email = ?"
	} else {
		c.JSON(http.StatusExpectationFailed, gin.H{"msg": "The field Email or Name do not be empty"})
		return
	}

	if err := config.DB.Where(Field, loginInput.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "There is no account with this user name."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for user"})
		}
		return
	}

	if !user.ComparePassword(loginInput.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Incorrect password"})
		return
	}

	token, err := utils.GenerateToken(user.Email, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": user, "Token": token})

}
