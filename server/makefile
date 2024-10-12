build:
	go build -o rk-server

gen-proto:
	protoc --go_out=. --go-grpc_out=. -I . kadev-rk.proto
