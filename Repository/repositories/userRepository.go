package repositories

/*
	- UserRepository for user
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

import (
	"github.com/jinzhu/gorm"
	"okky/go-design-pattern/Repository/models"
)

type UserRepository struct {
	DB        *gorm.DB
	TableName string
}

// initialize repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB:        db,
		TableName: "users",
	}
}

// example query get data by ID
func (repo *UserRepository) Get(id int) (models.User, error) {
	var user models.User
	query := repo.DB.Table(repo.TableName)
	query = query.Where("id = ?", id)
	query = query.Scan(&user)

	return user, query.Error
}