# Переменные
BINARY_NAME = goc
BUILD_DIR = build
GO_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKG = "./..."

## build: Собрать проект
build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/goc

## test: Запустить тесты
test:
	@go test ./...

## test-coverage: Запустить тесты с генерацией coverage отчета
test-coverage:
	@go test -coverprofile=$(BUILD_DIR)/coverage.out $(PKG) && go tool cover -html=$(BUILD_DIR)/coverage.out

## lint: Запустить линтер (golangci-lint)
lint:
	@if [ -x "$(shell command -v golangci-lint)" ]; then \
		golangci-lint run; \
	else \
		echo "golangci-lint не установлен. Установите: https://golangci-lint.run/"; \
	fi

## fmt: Форматировать код
fmt:
	@go fmt $(PKG)
	@if [ -x "$(shell command -v goimports)" ]; then \
		goimports -w .; \
	else \
		echo "goimports не установлен. Установите: go install golang.org/x/tools/cmd/goimports@latest"; \
	fi

## clean: Очистить сгенерированные файлы и папку build
clean:
	@go clean
	@rm -rf $(BUILD_DIR)

## deps: Проверить зависимости
deps:
	@go mod verify
	@go mod tidy

## install: Установить зависимости
install:
	@go mod download

## vendor: Создать vendor директорию
vendor:
	@go mod vendor

## add-to-path: Показать инструкции для добавления папки build в PATH
add-to-path:
	@mkdir -p $(BUILD_DIR)
	@echo "\nДобавьте папку с бинарником в PATH:\n"
	@build_path=$$(cd $(BUILD_DIR) && pwd); \
	if [ "$$(uname -s)" = "Darwin" ] || [ "$$(uname -s)" = "Linux" ]; then \
		echo "Для Linux/macOS (временное добавление):"; \
		echo "  export PATH=\$$PATH:$$build_path\n"; \
		echo "Для постоянного добавления (bash/zsh):"; \
		echo "  echo 'export PATH=\$$PATH:$$build_path' >> ~/.bashrc && source ~/.bashrc\n"; \
	elif [ "$$OS" = "Windows_NT" ]; then \
		win_path=$$(cygpath -w "$$build_path" 2>/dev/null || echo "$$build_path"); \
		echo "Для Windows (временное добавление):"; \
		echo "  set PATH=%%PATH%%;$$win_path"; \
		echo "\nДля постоянного добавления:"; \
		echo "  1. Правой кнопкой по 'Этот компьютер' → Свойства"; \
		echo "  2. Дополнительные параметры системы → Переменные среды"; \
		echo "  3. В разделе 'Системные переменные' выберите 'Path' → Изменить"; \
		echo "  4. Добавьте новую запись: $$win_path\n"; \
	else \
		echo "Неизвестная ОС. Путь к бинарнику: $$build_path"; \
	fi

## help: Показать доступные команды
help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

.PHONY: build test test-coverage lint fmt clean deps install vendor help add-to-path
