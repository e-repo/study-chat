include ./backend/.env
export

.PHONY: help unit_test integration_test test lint coverage_report cpu_profile mem_profile migrate_up migrate_down create_migration

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

unit_test:
	go test -v ./backend/internal/...

integration_test:
	go test -v ./backend/integration_test/...

test: unit_test integration_test

lint:
	go fmt ./backend/...
	find . -name '*.go' ! -path "./generated/*" -exec goimports -local study-chat/ -w {} +
	find . -name '*.go' ! -path "./generated/*" -exec golines -w {} -m 120 \;
	golangci-lint run ./...
	./check-go-generate.sh

coverage_report:
	# TODO: fix test execution in 1 thread
	go test -p=1 -coverpkg=./... -count=1 -coverprofile=.coverage.out ./...
	go tool cover -html .coverage.out -o .coverage.html
	open ./.coverage.html

cpu_profile:
	go test -cpuprofile=profiles/cpu.prof  ./e2e_test
	go tool pprof -http=:6061 profiles/cpu.prof

mem_profile:
	go test -memprofile=profiles/mem.prof ./e2e_test
	go tool pprof -http=:6061 profiles/mem.prof

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOSTS):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(if $(filter $(POSTGRES_SSL),true),require,disable)

migrate_up:
	migrate -path ./backend/migrations -database "$(DB_URL)" up

migrate_down:
	migrate -path ./backend/migrations -database "$(DB_URL)" down $(count)

create_migration:
	migrate create -ext sql -dir ./backend/migrations $(name)
