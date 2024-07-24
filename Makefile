DB_NAME=bookstore_n9
DB_USER=postgres
DB_PASSWORD=7479
DB_HOST=localhost
DB_PORT=5432


POSTGRES_DB=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate-up:
	migrate -path migrations -database $(POSTGRES_DB) up


migrate-down:
	migrate -path migrations -database $(POSTGRES_DB) down

