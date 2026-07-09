include .env
export

MIGRATIONS_DIR = DB/migrations
SEEDERS_DIR = DB/seeders

build:
	go build main.go
run:
	npx nodemon --exec "go run" ./main.go --signal SIGTERM

migrations_up:
	goose -dir db/migrations mysql "$(DBUSER):$(DBPASS)@$(DB_Net)($(DB_Addr))/$(DBName)" up

migrations_down:
	goose -dir db/migrations mysql "$(DBUSER):$(DBPASS)@$(DB_Net)($(DB_Addr))/$(DBName)" down

seed_up:
	goose -dir db/seeders mysql "$(DBUSER):$(DBPASS)@$(DB_Net)($(DB_Addr))/$(DBName)" up

seed_down:
	goose -dir db/seeders mysql "$(DBUSER):$(DBPASS)@$(DB_Net)($(DB_Addr))/$(DBName)" down

migration_create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

seeders_create:
	goose -dir $(SEEDERS_DIR) create $(name) sql