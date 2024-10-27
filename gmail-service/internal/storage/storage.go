package storage

import (
	"context"
	pb "gmail-service/internal/pkg/genproto"
)

type StorageI interface {
	User() UserI
	Draft() DraftI
	Outbox() OutboxI
	Inbox() InboxI
	Attachment() AttachmentI
}

type UserI interface {
	ChangeUserPFP(context.Context, *pb.UserChangePFPReq) (*pb.Void, error)
	ChangeUserPassword(context.Context, *pb.UserRecoverPasswordReq) (*pb.Void, error)
	ConfirmUser(context.Context, *pb.ByEmail) (*pb.Void, error)
	CreateUser(context.Context, *pb.UserCreateReq) (*pb.Void, error)
	DeleteUser(context.Context, *pb.ByID) (*pb.Void, error)
	GetUserByEmail(context.Context, *pb.ByEmail) (*pb.UserGetRes, error)
	GetUserByID(context.Context, *pb.ByID) (*pb.UserGetRes, error)
	GetUserSecurityByEmail(context.Context, *pb.ByEmail) (*pb.UserGetSecurityRes, error)
	IsEmailExists(context.Context, *pb.ByEmail) (*pb.UserEmailCheckRes, error)
	UpdateUser(context.Context, *pb.UserUpdateReq) (*pb.Void, error)
}

type DraftI interface {
	Create(context.Context, *pb.DraftCreateUpdateReq) (*pb.Void, error)
	Delete(context.Context, *pb.ByID) (*pb.Void, error)
	SendDraft(context.Context, *pb.ByID) (*pb.MessageSentRes, error)
	Update(context.Context, *pb.DraftCreateUpdateReq) (*pb.Void, error)
}

type OutboxI interface {
	ArchiveMessage(context.Context, *pb.ByID) (*pb.Void, error)
	Delete(context.Context, *pb.ByID) (*pb.Void, error)
	Get(context.Context, *pb.ByID) (*pb.OutboxMessageGetRes, error)
	GetAll(context.Context, *pb.OutboxMessagesGetAllReq) (*pb.OutboxMessagesGetAllRes, error)
	MoveToTrash(context.Context, *pb.ByID) (*pb.Void, error)
	Send(context.Context, *pb.OutboxMessageSentReq) (*pb.MessageSentRes, error)
	StarMessage(context.Context, *pb.ByID) (*pb.Void, error)
}

type InboxI interface {
	ArchiveMessage(context.Context, *pb.ByID) (*pb.Void, error)
	Delete(context.Context, *pb.ByID) (*pb.Void, error)
	GetAll(context.Context, *pb.InboxMessageGetAllReq) (*pb.InboxMessagesGetAllRes, error)
	GetByID(context.Context, *pb.ByID) (*pb.InboxMessageGetRes, error)
	MarkAsRead(context.Context, *pb.ByID) (*pb.Void, error)
	MarkAsSpam(context.Context, *pb.ByID) (*pb.Void, error)
	MoveToTrash(context.Context, *pb.ByID) (*pb.Void, error)
	StarMessage(context.Context, *pb.ByID) (*pb.Void, error)
}

type AttachmentI interface {
	Create(context.Context, *pb.AttachmentCreateReq) (*pb.AttachmentCreateRes, error)
	GetByID(context.Context, *pb.ByID) (*pb.AttachmentGetRes, error)
	Delete(context.Context, *pb.ByID) (*pb.AttachmentDeleteRes, error)
	GetAll(context.Context, *pb.AttachmentGetAllReq) (*pb.AttachmentGetAllRes, error)
	GetMyUploads(context.Context, *pb.ByID) (*pb.AttachmentGetAllRes, error)
}
