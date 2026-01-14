package main

import (
	"fff/internal/model"
	"fff/internal/repo"
	"fmt"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=chat_api port=5432 sslmode=disable"

	db, err := repo.InitDB(dsn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to database", db)

	chatRepo := repo.NewChatRepo(db)
	chatmodel := &model.Chat{Title: "Hello World"}
	err = chatRepo.Create(chatmodel)
	fmt.Println(chatmodel)
}
