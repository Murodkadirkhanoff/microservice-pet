package main

import (
	"log"
	"net"

	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/db"
	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/handlers"
	pb "github.com/Murodkadirkhanoff/pet-microservice-golang-app/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Запуск REST-сервера
	db.InitDB()
	db.MigrateDB()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &handlers.AuthServer{})

	reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port 50051")

	// Запускаем gRPC-сервер (блокирующая операция)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

}
