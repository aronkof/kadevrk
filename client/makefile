build:
	GOOS=windows go build -o rk-client.exe

gen-proto:
	protoc --go_out=. -I . kadev-rk.proto

clean-deps:
	del /f .\deps\gkl.exe

build-deps: clean-deps
	g++ .\deps\gkl.cpp -o .\deps\gkl.exe
