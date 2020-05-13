package book

import (
	"github.com/gofiber/fiber"
)

func Getbooks(c *fiber.Ctx){
	c.Send("Get all book")
}

func Getbook(c *fiber.Ctx){
	c.Send("Get a single book")
}

func Addbook(c *fiber.Ctx){
	c.Send("Add a book")
}

func Delelebook(c *fiber.Ctx){
	c.Send("Remove a book record")
}

func Updatebook(c *fiber.Ctx){
	c.Send("Update a book record")
}

