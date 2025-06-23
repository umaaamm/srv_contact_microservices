package contact

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"srv_contact/main/api/presenter"
	"srv_contact/main/pkg/entities"
)

type Repository interface {
	CreateContact(book *entities.Contact) (*entities.Contact, error)
	ReadContact() (*[]presenter.Contact, error)
	UpdateContact(book *entities.Contact) (*entities.Contact, error)
	DeleteContact(ID string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateContact(contact *entities.Contact) (*entities.Contact, error) {
	contact.ID = primitive.NewObjectID()
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *repository) ReadContact() (*[]presenter.Contact, error) {
	var contacts []presenter.Contact
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var contact presenter.Contact
		_ = cursor.Decode(&contact)
		contacts = append(contacts, contact)
	}
	return &contacts, nil
}

func (r *repository) UpdateContact(contact *entities.Contact) (*entities.Contact, error) {
	contact.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": contact.ID}, bson.M{"$set": contact})
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *repository) DeleteContact(ID string) error {
	contactID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": contactID})
	if err != nil {
		return err
	}
	return nil
}
