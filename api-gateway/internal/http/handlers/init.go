package handlers

import (
	"context"

	pb "api-gateway/internal/pkg/genproto"
	minio_connect "api-gateway/internal/pkg/minio"
	rdb "api-gateway/internal/pkg/redis"

	l "github.com/azizbek-qodirov/logger"
	"google.golang.org/grpc"
)

var (
	err error
)

type HTTPHandler struct {
	Logger *l.Logger
	RDB    *rdb.RedisClient
	Minio  *minio_connect.MinioClient
	US     pb.UserServiceClient
	DS     pb.DraftServiceClient
	IS     pb.InboxServiceClient
	OS     pb.OutboxServiceClient
	AS     pb.AttachmentServiceClient
}

func NewHandler(us *grpc.ClientConn, logger *l.Logger) *HTTPHandler {
	db, err := rdb.NewRedisClient(context.Background())
	if err != nil {
		logger.ERROR.Panicln("Redis not connected due to error: " + err.Error())
	}

	mc, err := minio_connect.NewMinioClient()
	if err != nil {
		logger.ERROR.Panicln("Minio not connected due to error: " + err.Error())
	}

	return &HTTPHandler{
		Logger: logger,
		RDB:    db,
		US:     pb.NewUserServiceClient(us),
		DS:     pb.NewDraftServiceClient(us),
		IS:     pb.NewInboxServiceClient(us),
		OS:     pb.NewOutboxServiceClient(us),
		AS:     pb.NewAttachmentServiceClient(us),
		Minio:  mc,
	}
}
