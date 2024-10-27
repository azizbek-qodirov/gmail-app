package app

import (
	"log"
	"net"

	"gmail-service/internal/pkg/config"
	pb "gmail-service/internal/pkg/genproto"
	"gmail-service/internal/pkg/postgres"
	"gmail-service/internal/storage/repo"
	"gmail-service/internal/usecase/service"

	"google.golang.org/grpc"
)

func Run(cf *config.Config) {
	pgm, err := postgres.New(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer pgm.Close()

	db := repo.NewStorage(pgm.DB)

	lis, err := net.Listen("tcp", cf.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, service.NewUserService(db))
	pb.RegisterDraftServiceServer(server, service.NewDraftService(db, db.DB))
	pb.RegisterInboxServiceServer(server, service.NewInboxService(db))
	pb.RegisterOutboxServiceServer(server, service.NewOutboxService(db, db.DB))
	pb.RegisterAttachmentServiceServer(server, service.NewAttachmentService(db))

	log.Println("Server started on port " + cf.GRPCPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer lis.Close()
}
