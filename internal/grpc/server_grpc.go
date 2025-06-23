package grpc

import (
	"srv_contact/main/pkg/contact"
	proto "srv_contact/main/proto/contact"
)

type server struct {
	proto.UnimplementedContactServiceServer
	repo contact.Repository
}

func NewGRPCServer(repo contact.Repository) proto.ContactServiceServer {
	return &server{repo: repo}
}
