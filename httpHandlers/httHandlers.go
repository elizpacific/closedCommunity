package httpHandlers

import (
	"net/http"

	"closedCommunity/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, db *gorm.DB) {
	var input struct {
		Nickname  string   `json:"nickname" binding:"required"`
		Email     string   `json:"email" binding:"required"`
		Phone     string   `json:"phone"`
		Position  string   `json:"position"`
		Graduated string   `json:"graduated"`
		Country   string   `json:"country"`
		City      string   `json:"city"`
		Bio       string   `json:"bio"`
		AboutMe   string   `json:"about_me"`
		FirstName string   `json:"first_name" binding:"required"`
		LastName  string   `json:"last_name" binding:"required"`
		Company   string   `json:"company"`
		Hobbies   []string `json:"hobbies"`
		Links     []string `json:"links"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Nickname:  input.Nickname,
		Email:     input.Email,
		Phone:     input.Phone,
		Position:  input.Position,
		Graduated: input.Graduated,
		Country:   input.Country,
		City:      input.City,
		Bio:       input.Bio,
		AboutMe:   input.AboutMe,
		Profile: models.UserProfile{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Company:   input.Company,
		},
	}

	// Додаємо хобі
	for _, h := range input.Hobbies {
		user.Hobbies = append(user.Hobbies, models.Hobby{Hobby: h})
	}

	// Додаємо посилання
	for _, l := range input.Links {
		user.Links = append(user.Links, models.Link{Link: l})
	}

	// Зберігаємо користувача
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}
