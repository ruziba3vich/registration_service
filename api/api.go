package api

import (
	"context"
	"log"
	"net"

	genprotos "github.com/ruziba3vich/registration_ms/genprotos/protos"
	"github.com/ruziba3vich/registration_ms/internal/storage"
	"google.golang.org/grpc"
)

type (
	API struct {
		// genprotos.UnimplementedMessageServiceServer
		genprotos.UnimplementedUserServiceServer
		storage *storage.Storage
	}
)

func New(s *storage.Storage) *API {
	return &API{
		storage: s,
	}
}

func (a *API) CreateUser(ctx context.Context, req *genprotos.CreateUserRequest) (*genprotos.CreateUserResponse, error) {
	return a.storage.CreateUser(context.Background(), req)
}

func (a *API) CreateAdmin(ctx context.Context, req *genprotos.CreateAdminRequest) (*genprotos.CreateAdminResponse, error) {
	return nil, nil
}

func (a *API) Run() error {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	creatorService := New(a.storage)
	genprotos.RegisterUserServiceServer(serverRegisterer, creatorService)

	if err := serverRegisterer.Serve(lis); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
