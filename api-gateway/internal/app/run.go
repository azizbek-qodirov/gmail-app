package app

import (
	api "api-gateway/internal/http"
	"api-gateway/internal/http/handlers"
	"api-gateway/internal/pkg/config"
	"context"
	"fmt"

	_ "api-gateway/internal/http/docs"

	rdb "api-gateway/internal/pkg/redis"

	l "github.com/azizbek-qodirov/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cf config.Config) {
	em := config.NewErrorManager()
	logger, err := l.NewLogger(&l.LogFileConfigs{
		Directory: "logs",
		Filename:  "app.log",
		Stdout:    false,
		Include:   l.DateTime | l.Loglevel | l.ShortFileName,
	})
	em.CheckErr(err)

	GmailServiceConn, err := grpc.NewClient(fmt.Sprintf("gmail_service%s", cf.GMAIL_GRPC_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer GmailServiceConn.Close()

	rdb, err := rdb.NewRedisClient(context.Background())
	if err != nil {
		logger.ERROR.Panicln("Redis not connected due to error: " + err.Error())
	}
	defer rdb.Close()

	handler := handlers.NewHandler(GmailServiceConn, logger)
	router := api.NewRouter(handler, rdb, logger)

	fmt.Println("Server is running on port:", cf.GATEWAY_HTTP_PORT)
	if err := router.Run(cf.GATEWAY_HTTP_PORT); err != nil {
		panic(err)
	}
}
