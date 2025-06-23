package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "srv_contact/main/proto/contact"
)

func (s *server) GetContactByID(ctx context.Context, req *proto.GetContactRequest) (*proto.ContactResponse, error) {
	contactID := req.Id

	data, err := s.repo.FindByID(contactID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "contact not found: %v", err)
	}

	return &proto.ContactResponse{
		Id:   data.ID.Hex(),
		Nama: data.Nama,
		NoHp: data.NoHp,
	}, nil
}
