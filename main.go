package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"srv_contact/main/api/router"
	"srv_contact/main/pkg/contact"
	pb "srv_contact/main/proto/contact"
)

type server struct {
	pb.UnimplementedContactServiceServer
}

func main() {
	// 1. Koneksi database
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatalf("Database Connection Error: %v", err)
	}
	defer cancel()
	fmt.Println("Database connection success!")

	// 2. Init repository dan service
	contactCollection := db.Collection("contacts")
	contactRepo := contact.NewRepo(contactCollection)
	contactService := contact.NewService(contactRepo)

	// 3. Jalankan REST API dengan Fiber (port :8080)
	go func() {
		app := fiber.New()
		app.Use(cors.New())
		app.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("Services Contact API is running")
		})
		api := app.Group("/api")
		router.ContactRouter(api, contactService)

		log.Println("REST API running on :8080")
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Fiber error: %v", err)
		}
	}()

	// 4. Jalankan gRPC server (port :50051)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterContactServiceServer(grpcServer, &server{})

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://username:password@localhost:27017/contact").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("contact")
	return db, cancel, nil
}
