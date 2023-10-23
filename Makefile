include internal/db/.env

LOCAL_BIN:=$(CURDIR)/bin
IMAGE_VERSION:=0.0.1
LOCAL_MIGRATION_DIR:=internal/db/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"


install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	

generate_user:
	make VERSION=1 API=user generate

generate:
	mkdir -p pkg/$(API)_v$(VERSION)
	protoc --proto_path api/$(API)_v$(VERSION) \
	--go_out=pkg/$(API)_v$(VERSION) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API)_v$(VERSION) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/$(API)_v$(VERSION)/$(API).proto

dockerBuildAndPush:
	docker buildx build --no-cache --platform linux/amd64 \
	-t cr.selcloud.ru/nazip-reestr/auth:$(IMAGE_VERSION) .
	docker login -u token -p CRgAAAAA20Et5rj42dc0m7h020YikYpsUtaadRgl cr.selcloud.ru/nazip-reestr
	docker push cr.selcloud.ru/nazip-reestr/auth:$(IMAGE_VERSION)

start-postgres:
	docker-compose -f internal/db/docker-compose.yml up -d

stop-postgres:
	docker-compose -f internal/db/docker-compose.yml down

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

gen-repo-users:
	$(LOCAL_BIN)/sqlc generate -f internal/repository/user/sqlc/user.yaml