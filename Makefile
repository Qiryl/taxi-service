protoc:
	@echo Generating user proto
	@cd proto/user && protoc --go_out=. --go-grpc_out=. user.proto
