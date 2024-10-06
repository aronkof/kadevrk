build:
	go build -o kadev-rk-server

gen-proto:
	protoc --go_out=. --go-grpc_out=. -I . kadev-rk.proto
