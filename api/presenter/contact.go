package presenter

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"srv_contact/main/pkg/entities"
)

type Contact struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nama string             `json:"nama"`
	NoHp string             `json:"noHp"`
}

func ContactSuccessResponse(data *entities.Contact) *fiber.Map {
	contact := Contact{
		ID:   data.ID,
		Nama: data.Nama,
		NoHp: data.NoHp,
	}
	return &fiber.Map{
		"status": true,
		"data":   contact,
		"error":  nil,
	}
}

func ContactsSuccessResponse(data *[]Contact) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ContactErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
