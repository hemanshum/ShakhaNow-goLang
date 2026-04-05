package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

var store = session.New()

// Simple in-memory user for the demo
type User struct {
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
}

var admin = &User{
	Name:     "Hemanshu Mahajan",
	Username: "hemanshu",
	Email:    "hemanshu.mahajan@gmail.com",
	Phone:    "1234567890",
	Password: "Password123",
}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	// Reload templates on each request in development
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views:          engine,
		ReadBufferSize: 32768, // Increase buffer size to handle large localhost cookies
	})

	// Add middlewares for beautiful logging and recovery
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: " [${time}] | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
	}))

	// Serve static files from the public folder (Tailwind CSS output)
	app.Static("/public", "./public")

	// Middleware to check authentication
	authMiddleware := func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Redirect("/login")
		}
		if sess.Get("authenticated") == nil {
			return c.Redirect("/login")
		}
		return c.Next()
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		if sess.Get("authenticated") != nil {
			return c.Redirect("/admin")
		}
		return c.Render("login", fiber.Map{
			"Error": c.Query("error"),
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		identifier := c.FormValue("email") // This field will accept email, username, or phone
		password := c.FormValue("password")

		// Check against email, username, or phone
		if (identifier == admin.Email || identifier == admin.Username || identifier == admin.Phone) && password == admin.Password {
			sess, err := store.Get(c)
			if err != nil {
				return c.Redirect("/login?error=System%20Error")
			}
			sess.Set("authenticated", true)
			sess.Set("email", admin.Email)
			sess.Save()
			return c.Redirect("/admin")
		}
		return c.Redirect("/login?error=Invalid%20Credentials")
	})

	app.Post("/logout", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		sess.Destroy()
		return c.Redirect("/login")
	})

	// Admin Routes Group
	adminGroup := app.Group("/admin", authMiddleware)

	// This matches exactly "/admin"
	adminGroup.Get("", func(c *fiber.Ctx) error {
		return c.Render("admin", fiber.Map{
			"Email": admin.Email,
			"Name":  admin.Name,
		})
	})

	// This matches "/admin/settings"
	adminGroup.Get("/settings", func(c *fiber.Ctx) error {
		return c.Render("settings", fiber.Map{
			"User":    admin,
			"Success": c.Query("success"),
			"Error":   c.Query("error"),
		})
	})

	adminGroup.Post("/settings", func(c *fiber.Ctx) error {
		currentPassword := c.FormValue("current_password")
		newPassword := c.FormValue("new_password")

		// Security Check: Current password MUST be provided and match
		if currentPassword != admin.Password {
			return c.Redirect("/admin/settings?error=Invalid%20Current%20Password")
		}

		// Update fields
		admin.Name = c.FormValue("name")
		admin.Username = c.FormValue("username")
		admin.Email = c.FormValue("email")
		admin.Phone = c.FormValue("phone")

		// Update password if a new one is provided
		if newPassword != "" {
			admin.Password = newPassword
		}

		return c.Redirect("/admin/settings?success=Profile%20Updated")
	})

	log.Println("Starting server on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
