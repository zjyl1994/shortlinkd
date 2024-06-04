package server

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/zjyl1994/shortlinkd/infra/vars"
)

var App *fiber.App

//go:embed index.html
var htmlTemplates embed.FS

func Run(listenAddr string) error {
	engine := html.NewFileSystem(http.FS(htmlTemplates), ".html")
	App = fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
	})
	App.Get("/", indexPage)
	App.Get("/:code", LinkHandler)
	if vars.DEBUG_MODE {
		debugG := App.Group("/debug")
		debugG.Get("/list", ListLinkHandler)
		debugG.Get("/reload", ReloadLinkHandler)
	}
	return App.Listen(listenAddr)
}

func indexPage(c *fiber.Ctx) error {
	if vars.INDEX_PAGE > "" {
		return c.Redirect(vars.INDEX_PAGE)
	}
	return c.Render("index", fiber.Map{})
}
