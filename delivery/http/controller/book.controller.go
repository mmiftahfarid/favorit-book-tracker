package controller

import (
	"favorit-book-tracker/delivery/http/dto/request"
	"favorit-book-tracker/domain"
	"favorit-book-tracker/shared/util"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	domain domain.Domain
}

func NewBookController(domain domain.Domain) BookController {
	return BookController{
		domain: domain,
	}
}

func (this *BookController) GetBook(ctx *fiber.Ctx) error {
	book, err := this.domain.BookUsecase.GetFavoritBook()

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch book.")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Render("resource/views/home.html", fiber.Map{
		"Title":  book.Title,
		"Author": book.Author,
		"Rating": book.Rating,
	})
}

func (this *BookController) AddFavoritBook(ctx *fiber.Ctx) error {
	var book request.RequestBookDTO
	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid rquest body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := this.domain.BookUsecase.SaveFavoritBook(book.ToBookEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}
