package main

import (
	"embed"
	"favorit-book-tracker/delivery/http"
	"favorit-book-tracker/domain"
	"fmt"
	"log"

	httpLib "net/http"
	"os"

	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var contentResource embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	domain := domain.ConstructionDomain()
	engine := html.NewFileSystem(httpLib.FS(contentResource), ".html")
	app := http.NewHttpDelivery(domain, engine)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
