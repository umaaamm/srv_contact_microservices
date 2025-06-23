package router

import (
	"github.com/gofiber/fiber/v2"

	"srv_contact/main/api/handlers"
	"srv_contact/main/pkg/contact"
)

func ContactRouter(app fiber.Router, service contact.Service) {
	app.Get("/contacts", handlers.GetContacts(service))
	app.Post("/contacts", handlers.AddContact(service))
	app.Put("/contacts", handlers.UpdateContact(service))
	app.Delete("/contacts", handlers.RemoveContact(service))
}
