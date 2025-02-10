package main

import (
	"log"
	"net"
	"services/name/internal/database"
	grpcserver "services/name/internal/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	// pb "services/auth/proto"
)

func main() {
	// Initialize database
	config := database.NewConfig()
	db, err := database.New(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Create gRPC server
	server := grpcserver.NewServer(db)
	grpcServer := grpc.NewServer()

	// Register services
	// pb.RegisterAuthServiceServer(grpcServer, server)
	grpc_health_v1.RegisterHealthServer(grpcServer, server)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}

	log.Printf("Starting gRPC server on :9090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
