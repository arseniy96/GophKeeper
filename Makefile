.PHONY: golangci-lint-run
golangci-lint-run: _golangci-lint-rm-unformatted-report

.PHONY: _golangci-lint-reports-mkdir
_golangci-lint-reports-mkdir:
	mkdir -p ./golangci-lint

.PHONY: _golangci-lint-run
_golangci-lint-run: _golangci-lint-reports-mkdir
	-docker run --rm \
    -v $(shell pwd):/app \
    -v $(shell pwd)/golangci-lint/cache:/root/.cache \
    -w /app \
    golangci/golangci-lint:v1.53.3 \
        golangci-lint run \
            -c .golangci.yml \
	> ./golangci-lint/report-unformatted.json

.PHONY: _golangci-lint-format-report
_golangci-lint-format-report: _golangci-lint-run
	cat ./golangci-lint/report-unformatted.json | jq > ./golangci-lint/report.json

.PHONY: _golangci-lint-rm-unformatted-report
_golangci-lint-rm-unformatted-report: _golangci-lint-format-report
	rm ./golangci-lint/report-unformatted.json

.PHONY: golangci-lint-clean
golangci-lint-clean:
	sudo rm -rf ./golangci-lint

POSTGRES_USER ?= gopher
POSTGRES_DB ?= goph_keeper
POSTGRES_PASSWORD ?= ps
CONTAINER_NAME ?= postgres-gophkeeper
IMAGE_NAME ?= postgres:latest

.PHONY: pg_docker_build
pg_docker_build:
	@echo "Pulling the latest PostgreSQL image..."
	docker pull $(IMAGE_NAME)

.PHONY: pg_docker_run
pg_docker_run:
	@echo "Starting PostgreSQL container..."
	docker run --name $(CONTAINER_NAME) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -d -p 5433:5432 $(IMAGE_NAME)

.PHONY: pg_docker_stop
pg_docker_stop:
	@echo "Stopping PostgreSQL container..."
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

.PHONY: pg_docker_clean
pg_docker_clean: pg_docker_stop
	@echo "Removing PostgreSQL image..."
	docker rmi $(IMAGE_NAME)