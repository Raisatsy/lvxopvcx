package repo

import (
	"fff/internal/model"

	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{db: db}
}

func (repo *ChatRepo) CreateChat(chat *model.Chat) error {
	return repo.db.Create(chat).Error
}

func (repo *ChatRepo) GetChatById(id uint, msgLimit int) (*model.Chat, error) {
	chat := model.Chat{}

	err := repo.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at ASC").Limit(msgLimit)
	}).First(&chat, id).Error

	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (repo *ChatRepo) CreateMessage(message *model.Message) error {
	return repo.db.Create(message).Error
}

func (repo *ChatRepo) DeleteChatById(id uint) error {
	return repo.db.Delete(&model.Chat{}, id).Error
}
