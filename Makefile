include ./backend/.env
export

.PHONY: help unit_test integration_test test lint coverage_report cpu_profile mem_profile migrate_up migrate_down create_migration

# frontend services
CHAT_NODE=node
NODE_CONTAINER_NAME=chat-node

help:
	cat Makefile

install:
	brew install protobuf
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/segmentio/golines@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

unit-test:
	go test -v ./backend/internal/...

integration-test:
	go test -v ./backend/integration_test/...

test: unit-test integration-test

lint:
	go fmt ./backend/...
	find . -name '*.go' ! -path "./generated/*" -exec goimports -local study-chat/ -w {} +
	find . -name '*.go' ! -path "./generated/*" -exec golines -w {} -m 120 \;
	golangci-lint run ./...
	./check-go-generate.sh

coverage-report:
	# TODO: fix test execution in 1 thread
	go test -p=1 -coverpkg=./... -count=1 -coverprofile=.coverage.out ./...
	go tool cover -html .coverage.out -o .coverage.html
	open ./.coverage.html

cpu-profile:
	go test -cpuprofile=profiles/cpu.prof  ./e2e_test
	go tool pprof -http=:6061 profiles/cpu.prof

mem-profile:
	go test -memprofile=profiles/mem.prof ./e2e_test
	go tool pprof -http=:6061 profiles/mem.prof

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOSTS):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(if $(filter $(POSTGRES_SSL),true),require,disable)

chat-init: docker-down chat-node-build chat-node-init


down: docker-down
ps: docker-ps
up: docker-up

ch-up: chat-node-up
ch-serve: chat-serve


docker-up:
	@docker compose up -d

chat-serve:
	@docker exec -it $(NODE_CONTAINER_NAME) yarn dev

chat-shell:
	@docker exec -it $(NODE_CONTAINER_NAME) sh

docker-ps:
	@docker compose ps

docker-down:
	docker compose down --remove-orphans

migrate-up:
	migrate -path ./backend/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./backend/migrations -database "$(DB_URL)" down $(count)

create-migration:
	migrate create -ext sql -dir ./backend/migrations $(name)

chat-node-up:
	docker compose up -d -- $(CHAT_NODE)

chat-node-build:
	docker compose build -- $(CHAT_NODE)

chat-node-init:
	docker compose run --rm $(CHAT_NODE) yarn install
