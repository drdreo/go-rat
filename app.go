package main

import (
"fmt"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)


// Field names should start with an uppercase letter
type Project struct {
    Name string `json:"name" xml:"name" form:"name"`
}

func main() {
    var projects []Project
	printSlice(projects)

    app := fiber.New()

    // Middleware
  	app.Use(recover.New())
  	app.Use(logger.New())

    // Match all routes starting with /api
    app.Use("/api", func(c *fiber.Ctx) error {
        fmt.Println("🥈 Second handler")
        return c.Next()
    })

    // GET /api/list
    app.Get("/api/projects", func(c *fiber.Ctx) error {
        return c.SendString("Projects!")
    })

    app.Post("/api/project", func(c *fiber.Ctx) error {

        project := new(Project)

        if err := c.BodyParser(project); err != nil {
            return err
        }

	    projects = append(projects, *project)
    	printSlice(projects)

        log.Println(project.Name)

        return c.Status(200).JSON(&fiber.Map{
            "success": true,
            "message": "",
            "data": project.Name,
        })
    })

    log.Fatal(app.Listen(*port))
}

func printSlice(projects []Project) {
	fmt.Printf("len=%d cap=%d %v\n", len(projects), cap(projects), projects)
}
