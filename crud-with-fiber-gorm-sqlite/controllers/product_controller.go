package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/services"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/viewmodels"
)

func GetAllProducts(c *fiber.Ctx) error {
	resultProducts := services.GetAllProducts()

	outputProducts := []viewmodels.Product{}
	for _, resultProduct := range resultProducts {
		outputProducts = append(outputProducts, viewmodels.Product{
			Id: resultProduct.Id,
			Name: resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		})
	}
	return c.Status(fiber.StatusOK).JSON(outputProducts)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid user id")
	}

	resultProduct, err := services.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	outputProduct := viewmodels.Product{
			Id: resultProduct.Id,
			Name: resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		}
	return c.Status(fiber.StatusOK).JSON(outputProduct)
}

func CreateProduct(c *fiber.Ctx) (error) {
	var inputProduct models.Product

	if err := c.BodyParser(&inputProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	
	services.CreateProduct(&inputProduct)

	outputProduct := viewmodels.Product{
			Id: inputProduct.Id,
			Name: inputProduct.Name,
			SerialNo: inputProduct.SerialNo,
		}

	return c.Status(fiber.StatusCreated).JSON(outputProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	// input
	product := models.Product{}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// validate & process
	resultProduct, err := services.UpdateProduct(int(product.Id), product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// output
	outputProduct := viewmodels.Product{
			Id: resultProduct.Id,
			Name: resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		}

	return c.Status(fiber.StatusCreated).JSON(outputProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	// input
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	
	// process
	resultProduct, err := services.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// output
	outputProduct := viewmodels.Product{
			Id: resultProduct.Id,
			Name: resultProduct.Name,
			SerialNo: resultProduct.SerialNo,
		}

	return c.Status(fiber.StatusCreated).JSON(outputProduct)

}
