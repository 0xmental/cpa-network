include .env
export

migrate-up:
	goose -dir ./db/migrations postgres "host=$(DB_HOST) user=$(POSTGRES_USER) dbname=$(POSTGRES_DB) password=$(POSTGRES_PASSWORD) sslmode=$(POSTGRES_SSL_MODE)" up

migrate-create:
	goose -dir db/migrations create $(name) sql
