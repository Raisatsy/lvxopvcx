package service

import (
	"errors"
	"fff/internal/model"
	"fff/internal/repo"
	"log/slog"
	"unicode/utf8"
)

type ChatService struct {
	repo *repo.ChatRepo
}

func NewChatService(repo *repo.ChatRepo) *ChatService {
	return &ChatService{repo: repo}
}

func (service *ChatService) CreateChat(title string) (*model.Chat, error) {
	count := utf8.RuneCountInString(title)
	if count < 1 || count > 200 {
		return nil, errors.New("Title must be between 1 and 200")
	}
	chatModel := model.Chat{
		Title: title,
	}

	if err := service.repo.CreateChat(&chatModel); err != nil {
		return nil, err
	}

	return &chatModel, nil
}

func (service *ChatService) AddMessageToChat(chatId uint, message string) (*model.Message, error) {
	count := utf8.RuneCountInString(message)
	if count < 1 || count > 5000 {
		return nil, errors.New("Message must be between 1 and 5000")
	}

	_, err := service.repo.GetChatById(chatId, 0)
	if err != nil {
		return nil, err
	}

	msg := model.Message{
		Text:   message,
		ChatID: chatId,
	}

	if err := service.repo.CreateMessage(&msg); err != nil {
		return nil, err
	}

	slog.Info("Successfully create message", "id", msg.ID, "message", msg.Text)
	return &msg, nil
}

func (service *ChatService) GetChatById(chatId uint, msgLimit int) (*model.Chat, error) {

	chatModel, err := service.repo.GetChatById(chatId, CheckLimit(msgLimit))

	if err != nil {
		return nil, err
	}

	slog.Info("Successfully got chat with messages", "id", chatModel.ID, "title", chatModel.Title)

	return chatModel, nil
}

func (service *ChatService) DeleteChatById(chatId uint) error {
	_, err := service.repo.GetChatById(chatId, 0)
	if err != nil {
		return err
	}

	return service.repo.DeleteChatById(chatId)
}

func CheckLimit(limit int) int {
	if limit > 100 {
		limit = 100
	}
	if limit < 1 {
		limit = 20
	}
	return limit
}
