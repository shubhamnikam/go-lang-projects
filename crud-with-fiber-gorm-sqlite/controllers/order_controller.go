package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/services"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/viewmodels"
)

func GetAllOrders(c *fiber.Ctx) error {
	//input
	//process
	resultOrders, err := services.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	//output
	outputOrders := []viewmodels.Order{}
	for _, resultOrder := range resultOrders {
		resultUser, err := services.GetUser(resultOrder.UserRefer)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		resultProduct, err := services.GetProduct(resultOrder.ProductRefer)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		outputOrders = append(outputOrders, viewmodels.Order{
			Id:        resultOrder.Id,
			CreatedAt: resultOrder.CreatedAt,
			User: viewmodels.User{
				Id:        resultUser.Id,
				FirstName: resultUser.FirstName,
				LastName:  resultUser.LastName,
			},
			Product: viewmodels.Product{
				Id:       resultProduct.Id,
				Name:     resultProduct.Name,
				SerialNo: resultProduct.SerialNo,
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(outputOrders)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid user id")
	}

	resultOrder, err := services.GetOrder(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resultUser, err := services.GetUser(resultOrder.UserRefer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resultProduct, err := services.GetProduct(resultOrder.ProductRefer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	outputOrder := viewmodels.Order{
		Id:        resultOrder.Id,
		CreatedAt: resultOrder.CreatedAt,
		User: viewmodels.User{
			Id:        resultUser.Id,
			FirstName: resultUser.FirstName,
			LastName:  resultUser.LastName,
		},
		Product: viewmodels.Product{
			Id:       resultProduct.Id,
			Name:     resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		},
	}
	return c.Status(fiber.StatusOK).JSON(outputOrder)
}

func CreateOrder(c *fiber.Ctx) error {
	var inputOrder models.Order

	if err := c.BodyParser(&inputOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	services.CreateOrder(&inputOrder)

	resultUser, err := services.GetUser(inputOrder.UserRefer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resultProduct, err := services.GetProduct(inputOrder.ProductRefer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	outputOrder := viewmodels.Order{
		Id:        inputOrder.Id,
		CreatedAt: inputOrder.CreatedAt,
		User: viewmodels.User{
			Id:        resultUser.Id,
			FirstName: resultUser.FirstName,
			LastName:  resultUser.LastName,
		},
		Product: viewmodels.Product{
			Id:       resultProduct.Id,
			Name:     resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(outputOrder)
}

func DeleteOrder(c *fiber.Ctx) error {
	// input
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// process
	resultOrder, err := services.DeleteOrder(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	resultUser, err := services.GetUser(resultOrder.UserRefer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resultProduct, err := services.GetProduct(resultOrder.ProductRefer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// output
	outputOrder := viewmodels.Order{
		Id:        resultOrder.Id,
		CreatedAt: resultOrder.CreatedAt,
		User: viewmodels.User{
			Id:        resultUser.Id,
			FirstName: resultUser.FirstName,
			LastName:  resultUser.LastName,
		},
		Product: viewmodels.Product{
			Id:       resultProduct.Id,
			Name:     resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(outputOrder)
}
