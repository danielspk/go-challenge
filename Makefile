ifneq (,$(wildcard ./deployments/.env))
    include deployments/.env
    export
endif

GO_APP=cmd/api/main.go
BINARY_OUTPUT=build/api

go-live:
	~/go/bin/air

go-clear:
	rm -f ${BINARY_OUTPUT}

go-fmt:
	go fmt ./...

go-test:
	go test ./...

go-build: go-clear
	go build -o ${BINARY_OUTPUT} ${GO_APP}

go-build-prod: go-clear go-fmt
	CGO_ENABLED=0 go build -a -ldflags "-s -w" -o ${BINARY_OUTPUT} ${GO_APP}

docker-build:
	docker-compose -f ./deployments/docker-compose.yml build

docker-up:
	docker-compose -f ./deployments/docker-compose.yml up

docker-down:
	docker-compose -f ./deployments/docker-compose.yml down --remove-orphans

docker-clear:
	sudo rm -rf ./deployments/data
