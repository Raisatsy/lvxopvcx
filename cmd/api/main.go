package main

import (
	"fff/internal/handler"
	"fff/internal/repo"
	"fff/internal/service"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=chat_api port=5432 sslmode=disable"
	}
	slog.Info("DSN", dsn)
	db, err := repo.InitDB(dsn)
	if err != nil {
		slog.Error("Error on connect to db", "err", err)
	}

	slog.Info("Connected to database")

	chatRepo := repo.NewChatRepo(db)
	chatService := service.NewChatService(chatRepo)
	chatHandler := handler.NewChatHandler(chatService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /chats", chatHandler.CreateChat)
	mux.HandleFunc("GET /chats/{id}", chatHandler.GetChat)
	mux.HandleFunc("POST /chats/{id}/messages", chatHandler.AddMessageToChat)
	mux.HandleFunc("DELETE /chats/{id}", chatHandler.DeleteChatById)

	host := "127.0.0.1:4000"

	slog.Info("Server started on", "host", host)
	if err := http.ListenAndServe(host, mux); err != nil {
		slog.Error("Server stopped", "Error", err)
	}
}
