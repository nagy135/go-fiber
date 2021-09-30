package main
import "github.com/gofiber/fiber/v2"
func main() {
  app := fiber.New()
  app.Get("/user/:name", func(c *fiber.Ctx) error {
      return c.SendString("Hello "+ c.Params("name"))
  })
  app.Listen(":8080")
}
