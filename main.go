package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	// "log"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 6433
	dbname   = "test_psql"
	user     = "admin"
	password = "password"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	print(psqlInfo)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	print(db)

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	fmt.Println("Connected Database")

	// err = createProduct(&Product{Name: "Go product 2", Price: 444})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Create Successful !")

	// product, err := getProduct(2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Get Successful", product)

	// product, err = updateProduct(6, &Product{Name: "UUU", Price: 333})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Update Successful !", product)

	// err = deleteProduct(8)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Delete Successful !")

	// products, err := getProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Get Successful !", products)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/product/:id", getProductHandler)

	app.Post("/product", createProductHandler)

	app.Put("/product/:id", updateProductHandler)

	app.Delete("/product/:id", deleteProductHandler)

	app.Get("/products", getProductsHandler)

	app.Listen(":9085")

}

func getProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	product, err := getProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	err := createProduct(p)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	return c.JSON(p)
}

func updateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	product, err := updateProduct(productId, p)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	err = deleteProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func getProductsHandler(c *fiber.Ctx) error {
	product, err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	return c.JSON(product)
}
