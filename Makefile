include .env
export


build:
	go build main.go
run:
	npx nodemon --exec "go run" ./main.go --signal SIGTERM

up:
	goose -dir db/migrations mysql "$(DBUSER):$(DBPASS)@$(DB_Net)($(DB_Addr))/$(DBName)" up
