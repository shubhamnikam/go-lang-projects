package services

import (
	"errors"

	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
)

func GetAllUsers() ([]models.User){
	users:= []models.User{}
	DbContext.Db.Find(&users)
	return users
}

func GetUser(id int) (models.User, error) {
	user := models.User{}
	DbContext.Db.Find(&user, "id = ?", id)
	if user.Id == 0 {
		return user, errors.New("User not found with id=" + string(id)) 
	} 

	return user, nil
}

func CreateUser(user *models.User) {
	DbContext.Db.Create(&user)
}

func UpdateUser(id int, inputUser models.User) (models.User, error) {
	//find 
	user := models.User{}
	DbContext.Db.Find(&user, "id = ?", id)
	if user.Id == 0 {
		return user, errors.New("User not found with id=" + string(id)) 
	} 
	//update
	user.FirstName = inputUser.FirstName
	user.LastName = inputUser.LastName
	DbContext.Db.Save(&user)
	if err := DbContext.Db.Save(&user).Error; err != nil {
		return user, errors.New("Error in update:"+err.Error())
	}

	return user, nil
}

func DeleteUser (id int) (models.User, error) {
	//find
	user := models.User{}
	DbContext.Db.Find(&user, "id=?", id)
	if user.Id == 0 {
		return user, errors.New("User not found with id="+string(id))
		} 
	//delete
	if err := DbContext.Db.Delete(&user).Error; err != nil {
		return user, errors.New("Error in deletion:"+err.Error())
	}
	
	return user, nil
}