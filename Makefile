current_dir = $(shell pwd)



BIN = bin
SERVER = api/



all:help

help:
	@echo run-server para executar o servidor
	@echo run-client para executar o cliente

#para dar setup
setup: 
	mkdir ../golib
	go env -w GOPATH=$(current_dir)/../golib:$(current_dir)
	make imports
	

build: 
	go build $(SERVER)controllers
	go build $(SERVER)auth
	go build $(SERVER)auto
	go build $(SERVER)middlewares
	go build $(SERVER)models
	go build $(SERVER)repository
	go build $(SERVER)responses
	go build $(SERVER)security
	go build config
	go build $(SERVER)database
	go build $(SERVER)router
	go build $(SERVER)

install-server:
	go install $(SERVER)

run-server:
	$(BIN)/server.exe

test: 
	go test $(SERVER)morestrings

run: build install-server run-server 

imports:
	go get github.com/gorilla/mux
	go get -u github.com/joho/godotenv
	go get github.com/jinzhu/gorm
	go get -u golang.org/x/crypto/bcrypt
	go get github.com/go-sql-driver/mysql
	go get github.com/dgrijalva/jwt-go
	go get github.com/gorilla/handlers
	

