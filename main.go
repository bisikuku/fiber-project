package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/go-fabric-http/book"
	"github.com/go-fabric-http/database"
)

func initialiseDb() {
	var err error
	database.DBconn, err = gorm.Open("sqlites", book.db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Open Database and ready to connect")
}
func main() {
	app := fiber.New()
	initialiseDb()
	defer database.DBconn.Close()
	routeBook(app)
	app.Listen(4000)

}

func routeBook(app *fiber.App) {
	app.Get("/api/v1/book", book.Getbooks)
	app.Get("/api/v1/book/:id", book.Getbook)
	app.Post("/api/v1/book", book.Addbook)
	app.Delete("/api/v1/book/:id", book.Delelebook)
	app.Patch("/api/v1/book/:id", book.Updatebook)
}
