package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/services"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/viewmodels"
)

func GetAllUsers(c *fiber.Ctx) error {
	resultUsers := services.GetAllUsers()

	outputUsers := []viewmodels.User{}
	for _, resultUser := range resultUsers {
		outputUsers = append(outputUsers, viewmodels.User{
			Id:        resultUser.Id,
			FirstName: resultUser.FirstName,
			LastName:  resultUser.LastName,
		})
	}
	return c.Status(fiber.StatusOK).JSON(outputUsers)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid user id")
	}

	resultUser, err := services.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	outputUser := viewmodels.User{
		Id:        resultUser.Id,
		FirstName: resultUser.FirstName,
		LastName:  resultUser.LastName,
	}
	return c.Status(fiber.StatusOK).JSON(outputUser)
}

func CreateUser(c *fiber.Ctx) error {
	var inputUser models.User

	if err := c.BodyParser(&inputUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	services.CreateUser(&inputUser)

	outputUser := viewmodels.User{
		Id:        inputUser.Id,
		FirstName: inputUser.FirstName,
		LastName:  inputUser.LastName,
	}

	return c.Status(fiber.StatusCreated).JSON(outputUser)
}

func UpdateUser(c *fiber.Ctx) error {
	// input
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// validate & process
	resultUser, err := services.UpdateUser(int(user.Id), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// output
	outputUser := viewmodels.User{
		Id:        resultUser.Id,
		FirstName: resultUser.FirstName,
		LastName:  resultUser.LastName,
	}

	return c.Status(fiber.StatusCreated).JSON(outputUser)
}

func DeleteUser(c *fiber.Ctx) error {
	// input
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// process
	resultUser, err := services.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// output
	outputUser := viewmodels.User{
		Id:        resultUser.Id,
		FirstName: resultUser.FirstName,
		LastName:  resultUser.LastName,
	}

	return c.Status(fiber.StatusCreated).JSON(outputUser)

}
