package api

import (
	"github.com/gin-gonic/gin"

	_ "api-gateway/internal/http/docs"
	"api-gateway/internal/http/handlers"
	"api-gateway/internal/http/middleware"
	rdb "api-gateway/internal/pkg/redis"

	l "github.com/azizbek-qodirov/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tag.name 01-Authentication
// @tag.description User registration, login, and email confirmation methods

// @tag.name 02-Password-recovery
// @tag.description Password recovery and reset functionality

// @tag.name 03-User-profile
// @tag.description User profile management including updates and deletion

// @tag.name 04-Inbox
// @tag.description Managing inbox messages such as reading, starring, and marking as spam

// @tag.name 05-Draft
// @tag.description Create, update, and send email drafts

// @tag.name 06-Outbox
// @tag.description Sending emails and managing sent items

// @tag.name 07-Attachments
// @tag.description Handling email attachments including uploading and deleting

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(h *handlers.HTTPHandler, rdb *rdb.RedisClient, logger *l.Logger) *gin.Engine {
	router := gin.Default()

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Authentication routes
	router.POST("/register", h.Register)
	router.POST("/confirm-registration", h.ConfirmRegistration)
	router.POST("/login", h.Login)
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/recover-password", h.RecoverPassword)
	router.POST("/send-code-again", h.SendCodeAgain)

	// Protected routes (require JWT authentication)
	protected := router.Group("/", middleware.JWTMiddleware(rdb))
	{
		// User-related routes
		protected.GET("/profile", h.Profile)
		protected.PUT("/user", h.UpdateProfile)
		protected.PUT("/user-password", h.UpdatePassword)
		protected.POST("/set-pfp", h.SetPFP)
		protected.DELETE("/delete-account", h.DeleteAccount)
		protected.GET("/user/:id", h.GetByID)

		// Draft routes
		protected.POST("/draft", h.CreateDraft)
		protected.PUT("/draft/:id", h.UpdateDraft)
		protected.DELETE("/draft/:id", h.DeleteDraft)
		protected.POST("/draft/:id/send", h.SendDraft)

		// Inbox routes
		protected.GET("/inbox", h.GetAllInboxMessages)
		protected.GET("/inbox/:id", h.GetInboxMessageByID)
		protected.PUT("/inbox/:id/trash", h.MoveInboxMessageToTrash)
		protected.DELETE("/inbox/:id", h.DeleteInboxMessage)
		protected.PUT("/inbox/:id/read", h.MarkInboxMessageAsRead)
		protected.PUT("/inbox/:id/spam", h.MarkInboxMessageAsSpam)
		protected.PUT("/inbox/:id/star", h.StarInboxMessage)
		protected.PUT("/inbox/:id/archive", h.ArchiveInboxMessage)

		// Outbox routes
		protected.POST("/outbox", h.SendMessage)
		protected.GET("/outbox", h.GetAllOutboxMessages)
		protected.GET("/outbox/:id", h.GetOutboxMessageByID)
		protected.PUT("/outbox/:id/trash", h.MoveOutboxMessageToTrash)
		protected.DELETE("/outbox/:id", h.DeleteOutboxMessage)
		protected.PUT("/outbox/:id/star", h.StarOutboxMessage)
		protected.PUT("/outbox/:id/archive", h.ArchiveOutboxMessage)

		// Attachment routes
		protected.POST("/attachment", h.CreateAttachment)
		protected.GET("/attachment/:outbox_id", h.GetAttachmentsByOutboxID)
		protected.DELETE("/attachment/:id", h.DeleteAttachment)
		protected.GET("/attachment/my-uploads", h.GetMyUploads)
	}

	return router
}
