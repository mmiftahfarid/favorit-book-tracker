package main

import (
	"embed"
	"favorit-book-tracker/domain"
	"html"

	"github.com/gofiber/template/html/v2"
	// _"fmt"
	httpLib "net/http"
	// _ "os"
	// "github.com/gofiber/template/html/v2"
)

var contentResource embed.FS

func main() {
	domain := domain.ConstructionDomain()
	engine := html.NewFileSystem(httpLib.FS(contentResource))
	// app := htt
}
