package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

const (
	db_host     = "host.docker.internal"
	db_port     = 5432
	db_user     = "dreo"
	db_password = "mypassword"
	db_name     = "go_rat"
)

const test_table_name = "project1"

// Field names should start with an uppercase letter
type Project struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func (p Project) String() string {
	return fmt.Sprintf("Project[%s]", p.Name)
}

func getProjectsString(ps []Project) string {
	var s string
	for _, p := range ps {
		s += p.String()
	}
	return s
}

func main() {
	app := fiber.New()

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", db_host, db_port),
		User:     db_user,
		Password: db_password,
		Database: db_name,
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful")

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
		var projects []Project
		err = db.Model(&projects).Select()
		if err != nil {
			panic(err)
		}

		return c.SendString(getProjectsString(projects))
	})

	app.Post("/api/project", func(c *fiber.Ctx) error {

		project := new(Project)

		if err := c.BodyParser(project); err != nil {
			return err
		}

		_, err = db.Model(project).Insert()
		if err != nil {
			panic(err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
			"data":    project.Name,
		})
	})

	log.Fatal(app.Listen(*port))
}

func printSlice(projects []Project) {
	fmt.Printf("len=%d cap=%d %v\n", len(projects), cap(projects), projects)
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Project)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
