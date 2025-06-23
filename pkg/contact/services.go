package contact

import (
	"srv_contact/main/api/presenter"
	"srv_contact/main/pkg/entities"
)

type Service interface {
	InsertContact(book *entities.Contact) (*entities.Contact, error)
	FetchContacts() (*[]presenter.Contact, error)
	UpdateContact(book *entities.Contact) (*entities.Contact, error)
	RemoveContact(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertContact(contact *entities.Contact) (*entities.Contact, error) {
	return s.repository.CreateContact(contact)
}

func (s *service) FetchContacts() (*[]presenter.Contact, error) {
	return s.repository.ReadContact()
}

func (s *service) UpdateContact(book *entities.Contact) (*entities.Contact, error) {
	return s.repository.UpdateContact(book)
}

func (s *service) RemoveContact(ID string) error {
	return s.repository.DeleteContact(ID)
}
