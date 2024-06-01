package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zjyl1994/shortlinkd/infra/vars"
)

func Run(listenAddr string) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Get("/:code", LinkHandler)
	if vars.DEBUG_MODE {
		debugG := app.Group("/debug")
		debugG.Get("/list", ListLinkHandler)
	}
	return app.Listen(listenAddr)
}
