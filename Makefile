include .env

export $(shell sed 's/=.*//' .env)
export ROOTDIR=$(shell pwd)

TEST_RUN ?= .
GOFLAGS ?= "-count=1"

run:
	@go run cmd/main.go

pb-regen:
	${HOME}/Apps/schemas/generate.sh ${HOME}/Apps/schemas/blogs ${ROOTDIR}/internal/pb

cleanTest:
	@go clean -testcache

test:
	@$(MAKE) cleanTest && go test -v ./tests/authorHelperTest/... -run ${TEST_RUN}
	@$(MAKE) cleanTest && go test -v ./tests/storeTest/... -run ${TEST_RUN}
	@$(MAKE) cleanTest && go test -v ./tests/storageTest/... -run ${TEST_RUN}
	@$(MAKE) cleanTest && go test -v ./tests/utilsTest/... -run ${TEST_RUN}

grpcTest:
	@$(MAKE) cleanTest && go test -v ./tests/grpcTest/... -run ${TEST_RUN}

gqlTest:
	@$(MAKE) cleanTest && go test -v ./tests/gqlTest/... -run ${TEST_RUN}

allTest:
	@${MAKE} cleanTest && go test -v ./tests/... -run ${TEST_RUN}