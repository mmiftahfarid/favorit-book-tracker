package router

import (
	"favorit-book-tracker/delivery/http/controller"
	"favorit-book-tracker/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	favoritBookController := controller.NewBookController(domain)

	app.Get("/", favoritBookController.GetBook)
	// app.Post("/", favoritBookController.GetBook)
}
