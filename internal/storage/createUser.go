package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	genprotos "github.com/ruziba3vich/registration_ms/genprotos/protos"
	"github.com/ruziba3vich/registration_ms/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	Storage struct {
		db *sql.DB
		// conn   *websocket.Conn
		config *config.Config
	}
)

func NewStorage(config *config.Config) *Storage {
	db, err := ConnectDB(*config)
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
		db: db,
		// conn:   conn,
		config: config,
	}
}

func (s *Storage) CreateUser(ctx context.Context, req *genprotos.CreateUserRequest) (*genprotos.CreateUserResponse, error) {
	query := `
		INSERT INTO users (
			username,
			data
		)
		VALUES ($1, $2)
		RETURNING user_id, username, data;
	`
	row := s.db.QueryRowContext(ctx, query, req.Username, req.Data)
	var response genprotos.CreateUserResponse
	if err := row.Scan(&response.UserId, &response.Username, &response.Data); err != nil {
		return nil, err
	}

	adminIds, err := s.getAllAdminUsernames(&ctx)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("user has been created with id : %s", response.UserId)

	for i := range adminIds {
		s.sendMessage(response.UserId, adminIds[i], message)
	}
	return &response, nil
}

func (s *Storage) getAllAdminUsernames(ctx *context.Context) ([]string, error) {
	query := `
		SELECT admin_id FROM admins;
	`
	var response []string
	rows, err := s.db.QueryContext(*ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resp string
		if err := rows.Scan(&resp); err != nil {
			return nil, err
		}
		response = append(response, resp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Storage) sendMessage(from, to, message string) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := genprotos.NewMessageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMessage(ctx, &genprotos.MessageRequest{From: from, To: to, Message: message})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("Message status: %s", r.Status)
}
