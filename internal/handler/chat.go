package handler

import (
	"encoding/json"
	"fff/internal/service"
	"log/slog"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	service *service.ChatService
}

func NewChatHandler(service *service.ChatService) *ChatHandler {
	return &ChatHandler{service: service}
}

func (ch *ChatHandler) CreateChat(res http.ResponseWriter, req *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		slog.Error("Failed to decode payload chat data", "error", err)
		http.Error(res, "Invalid payload data", http.StatusBadRequest)
		return
	}

	chatModel, err := ch.service.CreateChat(input.Title)

	if err != nil {
		slog.Error("Failed to create chat", "error", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info("Successfully created chat", "id", chatModel.ID, "title", chatModel.Title)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(chatModel)
}

func (ch *ChatHandler) GetChat(res http.ResponseWriter, req *http.Request) {
	chatId, err := strconv.ParseUint(req.PathValue("id"), 10, 32)

	if err != nil {
		slog.Error("Failed to parse chat id", "error", err)
		http.Error(res, "Invalid id", http.StatusBadRequest)
		return
	}

	limitStr := req.URL.Query().Get("limit")
	limit := 0
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			slog.Warn("Invalid limit format", "value", limitStr)
			http.Error(res, "Limit must be a number", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}
	chatModel, err := ch.service.GetChatById(uint(chatId), limit)

	if err != nil {
		slog.Error("Failed to get chat", "error", err)
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(chatModel)
}

func (ch *ChatHandler) AddMessageToChat(res http.ResponseWriter, req *http.Request) {
	chatId, err := strconv.ParseUint(req.PathValue("id"), 10, 32)
	if err != nil {
		slog.Error("Failed to parse chat id", "error", err)
		http.Error(res, "Invalid id", http.StatusBadRequest)
		return
	}

	var input struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		slog.Error("Failed to decode payload chat data", "error", err)
		http.Error(res, "Invalid payload data", http.StatusBadRequest)
		return
	}
	msg, err := ch.service.AddMessageToChat(uint(chatId), input.Text)

	if err != nil {
		slog.Error("Failed to create message", "error", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(msg)

}

func (ch *ChatHandler) DeleteChatById(res http.ResponseWriter, req *http.Request) {
	chatId, err := strconv.ParseUint(req.PathValue("id"), 10, 32)
	if err != nil {
		slog.Info("Cant parse id", "id", chatId, "err", err)
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = ch.service.DeleteChatById(uint(chatId))
	if err != nil {
		slog.Info("Cant delete chat", "id", chatId, "err", err)
		http.Error(res, "Chat not found", http.StatusNotFound)
		return
	}

	slog.Info("Successfully deleted chat", "id", chatId)

	res.WriteHeader(http.StatusNoContent)
}
