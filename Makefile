include internal/db/.env

LOCAL_BIN:=$(CURDIR)/bin
IMAGE_VERSION:=0.0.2
LOCAL_MIGRATION_DIR:=internal/db/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
		@if [ ! -f $(LOCAL_BIN)/protoc-gen-go-grpc ]; then \
			GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2;\
		fi
		@if [ ! -f $(LOCAL_BIN)/protoc-gen-go ]; then \
			GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1;\
		fi
		@if [ ! -f $(LOCAL_BIN)/goose ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0;\
		fi
		@if [ ! -f $(LOCAL_BIN)/sqlc ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest;\
		fi
		@if [ ! -f $(LOCAL_BIN)/minimock ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest;\
		fi
		@if [ ! -f $(LOCAL_BIN)/protoc-gen-openapiv2 ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2;\
		fi
		@if [ ! -f $(LOCAL_BIN)/protoc-gen-grpc-gateway ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2;\
		fi
		@if [ ! -f $(LOCAL_BIN)/statik ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7;\
		fi
		@if [ ! -f $(LOCAL_BIN)/golangci-lint ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3;\
		fi
		@if [ ! -f $(LOCAL_BIN)/protoc-gen-validate ]; then \
			GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1;\
		fi

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

test:
	go test ./...

generate-swagger:
	mkdir -p pkg/swagger
	make generate_user
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

generate_user_api:
	make VERSION=1 API=user generate

generate_auth_api:
	make VERSION=1 API=auth generate

generate_access_api:
	make VERSION=1 API=access generate

generate:
	mkdir -p pkg/$(API)_v$(VERSION)
	protoc --proto_path api/$(API)_v$(VERSION) --proto_path vendor.protogen \
	--go_out=pkg/$(API)_v$(VERSION) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API)_v$(VERSION) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/$(API)_v$(VERSION) --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/$(API)_v$(VERSION) --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
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

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi