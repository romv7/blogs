include .env

export $(shell sed 's/=.*//' .env)
export ROOTDIR=$(shell pwd)

run:
	@go run cmd/main.go

pb-regen:
	${HOME}/Apps/schemas/generate.sh ${HOME}/Apps/schemas/blogs ${ROOTDIR}/internal/pb

test:
	@go test -v ./tests/... -count=1
