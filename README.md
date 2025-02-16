# Учебный проект (study-chat)

### Линтеры и кодогенерация

Для запуска `make lint` и `go generate ./...` необходимо установить следующие утилиты:

```sh
brew install protobuf
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/segmentio/golines@latest
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

либо просто выполнить команду `make install`

### Pre-commit hooks

Для установки pre-commit hooks выполните команды:

```sh
brew install pre-commit
pre-commit install
```

### Golang-migrate

Для выполнения миграций базы данных используется утилита `golang-migrate`.
Чтобы установить утилиту, выполните команду:

```shell
brew install golang-migrate
```

Чтобы применить миграции к базе, существует make команда:

```shell
make migrate_up
```

Также для отката миграций (параметр `count` указывает количество миграций, которые нужно откатить):
```shell
make migrate_down count=1
```

Для создания новой миграции используйте команду:
```shell
make create_migration name=migration_name
```
