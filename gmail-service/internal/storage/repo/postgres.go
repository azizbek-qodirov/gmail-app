package repo

import (
	"database/sql"

	"gmail-service/internal/storage"
)

type Storage struct {
	UserS       storage.UserI
	DraftS      storage.DraftI
	OutboxS     storage.OutboxI
	InboxS      storage.InboxI
	AttachmentS storage.AttachmentI
	DB          *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		UserS:       NewUserRepo(db),
		AttachmentS: NewAttachmentRepo(db),
		DraftS:      NewDraftRepo(db),
		OutboxS:     NewOutboxRepo(db),
		InboxS:      NewInboxRepo(db),
		DB:          db,
	}
}

func (s *Storage) User() storage.UserI {
	return s.UserS
}

func (s *Storage) Draft() storage.DraftI {
	return s.DraftS
}

func (s *Storage) Outbox() storage.OutboxI {
	return s.OutboxS
}

func (s *Storage) Inbox() storage.InboxI {
	return s.InboxS
}

func (s *Storage) Attachment() storage.AttachmentI {
	return s.AttachmentS
}
