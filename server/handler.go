package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zjyl1994/shortlinkd/infra/vars"
	"github.com/zjyl1994/shortlinkd/service/code"
)

func LinkHandler(c *fiber.Ctx) error {
	shortCode := c.Params("code")
	config := code.GetCode(shortCode)
	if config == nil {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}
	if !config.Enabled {
		return c.Status(fiber.StatusNotFound).SendString("Disabled")
	}
	if config.Expired != nil && time.Now().After(*config.Expired) {
		return c.Status(fiber.StatusNotFound).SendString("Expired")
	}
	return c.Redirect(config.URL, fiber.StatusFound)
}

func ListLinkHandler(c *fiber.Ctx) error {
	items := code.ListCodes()
	return c.JSON(items)
}

func ReloadLinkHandler(c *fiber.Ctx) error {
	cfg, err := vars.LoadConfig(vars.CONFIG_FILE)
	if err != nil {
		return err
	}
	err = vars.ApplyConfig(cfg)
	if err != nil {
		return err
	}
	return c.SendString("OK")
}
