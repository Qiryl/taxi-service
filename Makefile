protoc:
	@echo Generating user proto
	@cd proto/user && protoc --go_out=. --go-grpc_out=. user.proto

build: build.client build.server

build.client:
	go build -o ./bin/client ./internal/user/cmd/client/client.go

build.server:
	go build -o ./bin/server ./internal/user/cmd/server/server.go

run: run_client run_server

run_client:
	@./bin/client &

run_server:
	@./bin/server &

dev.up:
	docker-compose -f docker-compose.dev.yml up 

dev.up.build:
	docker-compose -f docker-compose.dev.yml up --build

dev.down:
	docker-compose -f docker-compose.dev.yml down 
