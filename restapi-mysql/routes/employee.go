package routes

import (
	"github.com/WahidinAji/fiber-example/restapi-mysql/controllers"
	"github.com/gofiber/fiber/v2"
)

func employee(app *fiber.App){
	api := app.Group("/api")
	api.Get("/employees", controllers.GetAll)
	api.Get("/employees/:id", controllers.GetById)
	api.Post("/employees", controllers.CreateEmployee)
	api.Put("/employees/:id", controllers.UpdateEmployee)
	api.Delete("/employees/:id", controllers.DeleteEmployee)
}