.DEFAULT_GOAL := build

.PHONE:fmt vet build
fmt:
	go fmt ./...

vet: fmt
		go vet ./...

build: vet
		go build
