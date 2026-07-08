build:
	go build main.go
run:
	npx nodemon --exec "go run" ./main.go --signal SIGTERM