local:
	go run cmd/taylor-cli/main.go api -c configs/local.yaml

dev:
	go run cmd/taylor-cli/main.go api -c configs/dev.yaml

test:
	go run cmd/taylor-cli/main.go test -c configs/dev.yaml

build:
	GOOS=linux GOARCH=amd64 go build -o bin/taylor-cli cmd/taylor-cli/main.go
