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

func (repo *ChatRepo) Create(chat *model.Chat) error {
	return repo.db.Create(chat).Error
}

func (repo *ChatRepo) GetById(id uint) (*model.Chat, error) {
	chat := model.Chat{}
	err := repo.db.First(&chat, id).Error
	if err != nil {
		return nil, err
	}
	return &chat, nil
}
