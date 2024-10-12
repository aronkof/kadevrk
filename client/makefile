build:
	GOOS=windows go build -o kadev-rk-client-v2.exe

gen-proto:
	protoc --go_out=. --go-grpc_out=. -I . kadev-rk.proto

clean-deps:
	del /f .\deps\gkl.exe

build-deps: clean-deps
	g++ .\deps\gkl.cpp -o .\deps\gkl.exe
