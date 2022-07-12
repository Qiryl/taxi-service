protoc:
	@echo Generating user proto
	@cd proto/user && protoc --go_out=. --go-grpc_out=. user.proto

dev.up:
	docker-compose -f docker-compose.dev.yml up 

dev.up.build:
	docker-compose -f docker-compose.dev.yml up --build

dev.down:
	docker-compose -f docker-compose.dev.yml down 
