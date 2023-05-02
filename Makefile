export ROOT_DIR=$(shell pwd)

migrate:
	go run cmd/main.go

pb-regen:
	${HOME}/Apps/schemas/generate.sh ${HOME}/Apps/schemas/blogs ${ROOT_DIR}/pb

test:
	go test -v ./...
