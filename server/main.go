package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spammeturret/go-react-todo/pkg"
	"github.com/spammeturret/go-react-todo/storage"
)

func main() {
	tempPath := "../data-store/todos.json"
	fmt.Print("hello world")

	app := fiber.New()
	//app.use adds middleware to the fiber application
	//cors.New(cors.config) creates a new instance of the CORS middleware layer
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []pkg.Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &pkg.Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)
		storage.SaveTodoToFile(tempPath, todos)

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)

	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		// tempPath := "/Users/hiny/Documents/development/react/todo-application/data-store/todos.json"

		response, err := storage.LoadTodoFromJson(tempPath)
		if err == nil {
			todos = response
			return c.JSON(response)
		} else {
			return err
		}
	})
	log.Fatal(app.Listen(":4000"))
}
