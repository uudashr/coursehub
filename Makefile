DB_USER ?= coursehub
DB_PASSWORD ?= coursehubsecret
DB_PORT ?= 3306
DB_ADDRESS ?= 127.0.0.1:${DB_PORT}
DB_NAME ?= coursehub_test

# Dependendency management
.PHONY: vendor-prepare
vendor-prepare:
	@echo 'Install "dep"'
	@go get -u github.com/golang/dep/cmd/dep

.PHONY: vendor-update
vendor-update:
	@dep ensure -update

vendor:
	@dep ensure -vendor-only

# MySQL
.PHONY: docker-mysql-up
docker-mysql-up:
	@docker run --rm -d --name coursehub-mysql -p ${DB_PORT}:3306 -e MYSQL_DATABASE=$(DB_NAME) -e MYSQL_USER=$(DB_USER) -e MYSQL_PASSWORD=$(DB_PASSWORD) -e MYSQL_ROOT_PASSWORD=rootsecret mysql:5 && docker logs -f coursehub-mysql

.PHONY: docker-mysql-down
docker-mysql-down:
	@docker stop coursehub-mysql

# Mockery
.PHONY: mockery-prepare
mockery-prepare:
	go get github.com/vektra/mockery/.../

create-mock:
	mockery -name=Repository -dir=internal/account -output=internal/account/mocks

# Test
.PHONY: test
test: vendor 					# unit test
	@go test $(TEST_OPTS) ./...

.PHONY: test-mysql
test-mysql: vendor 				# mysql test
	@go test -tags=integration $(TEST_OPTS) ./internal/mysql/...

.PHONY: test-all
test-all: vendor 				# all test (include integration)
	@go test -tags=integration $(TEST_OPTS) ./...

# Database Migration
.PHONY: migrate-prepare
migrate-prepare:
	@echo 'Install "migrate"'
	@go get -u -d github.com/golang-migrate/migrate/cli github.com/go-sql-driver/mysql
	@cd $(shell go env GOPATH)/src/github.com/golang-migrate/migrate/cli
	@dep ensure
	@go build -tags 'mysql' -o /usr/local/bin/migrate github.com/golang-migrate/migrate/cli

# Tools
.PHONY: prepare-all
prepare-all: vendor-prepare mockery-prepare migrate-prepare
