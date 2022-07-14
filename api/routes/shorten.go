package routes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/zephyrus21/rego-short/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL                 string        `json:"url"`
	CustomShort         string        `json:"short"`
	Expiry              time.Duration `json:"expiry"`
	XRateLimitRemaining int           `json:"rate_limit"`
	XRateLimitRest      time.Duration `json:"rate_limit_rest"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	//? implement rate limiting

	//? check if the input URL is valid
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	//? check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Domain error",
		})
	}

	//? enforce https
	body.URL = helpers.EnforceHTTP(body.URL)

	return nil
}
