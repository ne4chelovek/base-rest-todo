include .env
export

LOCAL_BIN := $(CURDIR)/bin
LOCAL_MIGRATION_DIR ?= $(MIGRATION_DIR)
LOCAL_MIGRATION_DSN ?= "user=$(PG_USER) password=$(PG_PASSWORD) dbname=$(PG_DATABASE_NAME) host=localhost port=$(PG_PORT) sslmode=disable"



install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0


local-migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v


