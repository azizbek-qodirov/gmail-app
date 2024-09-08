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
	// connect to postgres
	pgm, err := postgres.New(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer pgm.Close()

	// connect to kafka producer
	// kf_p, err := kafka.NewKafkaProducer([]string{cf.KafkaUrl})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// connect to minio
	// minio, err := minio.MinIOConnect(cf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// repo
	db := repo.NewStorage(pgm.DB)
	// register kafka handlers
	// k_handler := KafkaHandler{
	// 	certificate: service.NewCertificateService(db, kf_p, minio),
	// }

	// if err := Registries(&k_handler, cf); err != nil {
	// 	log.Fatal(err)
	// }

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

	log.Println("Server started on port: ", cf.GRPCPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer lis.Close()
}

// func Registries(k_handler *KafkaHandler, cfg *config.Config) error {
// 	brokers := []string{cfg.KafkaUrl}
// 	kcm := kafka.NewKafkaConsumerManager()

// 	if err := kcm.RegisterConsumer(brokers, "certificate-update", "certificate-u", k_handler.CertificateUpdate()); err != nil {
// 		if err == kafka.ErrConsumerAlreadyExists {
// 			return errors.New("consumer for topic 'certificate-update' already exists")
// 		} else {
// 			return errors.New("error registering consumer:" + err.Error())
// 		}
// 	}
// 	return nil
// }
