.PHONY: help up down build-go test-go test-node test lint-go lint-node lint clean setup proto

# ============================================================
# AutoAPI-AK-Core — Root Makefile
# ============================================================

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Первоначальная настройка
	@cp -n .env.example .env || true
	@echo "Fill in .env before running"

up: ## Поднять все сервисы
	docker compose -f docker-compose.dev.yml up --build

down: ## Остановить все сервисы
	docker compose -f docker-compose.dev.yml down

build-go: ## Собрать Go сервисы
	$(MAKE) -C services/router build
	$(MAKE) -C services/context-bridge build

test-go: ## Тесты Go
	$(MAKE) -C services/router test
	$(MAKE) -C services/context-bridge test

test-node: ## Тесты Node
	cd services/gateway && npm test
	cd services/account-vault && npm test
	cd services/auth && npm test

test: test-go test-node ## Все тесты

lint-go: ## Линтинг Go
	$(MAKE) -C services/router lint
	$(MAKE) -C services/context-bridge lint

lint-node: ## Линтинг Node
	cd services/gateway && npm run lint
	cd services/account-vault && npm run lint
	cd services/auth && npm run lint

lint: lint-go lint-node ## Весь линтинг

clean: ## Очистить артефакты
	find services/router services/context-bridge -name bin -type d -exec rm -rf {} + 2>/dev/null || true
	find services -name dist -type d -exec rm -rf {} + 2>/dev/null || true

proto: ## Генерация из .proto
	bash scripts/generate-proto.sh
