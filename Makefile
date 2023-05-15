include .env

export $(shell sed 's/=.*//' .env)
export ROOTDIR=$(shell pwd)

TEST_RUN ?= .

run:
	@go run cmd/main.go

pb-regen:
	${HOME}/Apps/schemas/generate.sh ${HOME}/Apps/schemas/blogs ${ROOTDIR}/internal/pb

test:
	go clean -testcache && GOFLAGS="-count=1" go test -v ./tests/... -run ${TEST_RUN}
