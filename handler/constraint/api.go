package constraint

import (
	"eureka/handler"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetConstraints(c *fiber.Ctx) error {
	//Pulls constraints from DB
	constraintJson, err := GetConstraintsAsGeojson(c.Query("name"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid query param"})
	}
	//If request a viewer, forward data to viewer
	if strings.Contains(string(c.Request().URI().QueryString()), "view") {
		return handler.ViewInGeojsonio(c, constraintJson)
	}
	//Else return as a geojson
	return c.JSON(constraintJson)
}
