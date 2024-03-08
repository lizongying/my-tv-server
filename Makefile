.PHONY: server

all:  server

server:
	go vet ./cmd/server
	go build -ldflags "-s -w" -o  ./releases/my-tv-server  ./cmd/server

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/my-tv-server_linux_amd64 ./cmd/server

	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/my-tv-server_linux_arm64 ./cmd/server

	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/my-tv-server_darwin_amd64 ./cmd/server

	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/my-tv-server_darwin_arm64 ./cmd/server

	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/my-tv-server_windows_amd64.exe ./cmd/server