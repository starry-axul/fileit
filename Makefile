.PHONY: build clean deploy
export GOSUMDB=off

install:
	git config --local core.hooksPath .githooks/
	export GOSUMDB=off
	npm i
	git config url."git@gitlab.com:".insteadOf "https://gitlab.com/"
	go mod tidy
	go mod download

.PHONY: format
format:
	@go fmt ./internal/... ./pkg/... ./cmd/...

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/client-getall 	cmd/client/getall/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/client-create 	cmd/client/create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/client-regen 	cmd/client/regen/main.go

start:
	make format
	make build
	sls offline --useDocker

dev:
	docker-compose up -d
	make build
	sls offline --printOutput

deploy:
	make build
	sls deploy

test-cover:
	go test ./internal/... ./pkg/... -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

test:
	go test ./internal/... ./pkg/... -v

