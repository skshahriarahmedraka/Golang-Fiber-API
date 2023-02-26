package main

import "github.com/gofiber/fiber/v2"

func init(){
	
}

func main() {
	// Custom config
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1",
})

// ...

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/logout",func (c *fiber.Ctx) error {
		// Clears all cookies:
		c.ClearCookie()
		return c.SendString("Logout")
	})
	

	app.Listen(":8080")
}