all: web

web:
	go build -o ./bin/app ./main.go
