PKG := github.com/timPrachasri/app-chat/server
PKG_LIST := $(shell go list ${PKG}/...)

dev:
	@echo "============= starting locally ============="
	docker-compose -f resources/docker/docker-compose.yaml up --build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

down:
	@echo "============= downing core service ============="
	docker-compose -f resources/docker/docker-compose.yaml down

test:
	@go test -v ${PKG_LIST}

migrate:
	migrate create -ext sql -dir app/migrations $(name)

format:
	go fmt ./app/...

mockery_module:
	mockery -dir ./app/modules/$(module) -name $(interface) -output ./app/modules/$(module)/mocks

mockery:
	mockery -dir $(module_dir) -name $(interface) -output $(module_dir)/mocks