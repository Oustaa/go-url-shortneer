package services

import (
	"log"

	"github.com/oustaa/go-url-shortner/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

type UserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (us UserService) CreateUser(body UserPayload) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hashString := string(hashedPassword)

	user := models.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: hashString,
	}

	result := us.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (us UserService) GetUserByLogin(login string) (*models.User, error) {
	var user models.User

	result := us.db.
		Where("username LIKE ? OR email LIKE ?", "%"+login+"%", "%"+login+"%").
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
