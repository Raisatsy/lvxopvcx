DB_DRIVER=postgres
DB_DSN="host=127.0.0.1 user=postgres password=postgres dbname=chat_api sslmode=disable"
MIGRATIONS_DIR=./migrations

migration:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_DSN) up

down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_DSN) down

status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_DSN) status