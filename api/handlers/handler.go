package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"srv_contact/main/api/presenter"
	"srv_contact/main/pkg/contact"
	"srv_contact/main/pkg/entities"
)

func AddContact(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Contact
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		if requestBody.Nama == "" || requestBody.NoHp == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(errors.New(
				"Please specify name and noHp")))
		}
		result, err := service.InsertContact(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		return c.JSON(presenter.ContactSuccessResponse(result))
	}
}

func UpdateContact(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Contact
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		result, err := service.UpdateContact(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		return c.JSON(presenter.ContactSuccessResponse(result))
	}
}

func RemoveContact(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		contactID := requestBody.ID
		err = service.RemoveContact(contactID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetContacts(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchContacts()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}
		return c.JSON(presenter.ContactsSuccessResponse(fetched))
	}
}
