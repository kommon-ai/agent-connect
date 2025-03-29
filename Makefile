.PHONY: all generate run-server run-client clean setup-node setup-ts ts-build ts-client

all: generate

# Protoファイルからコードを生成
generate:
	buf generate

# Node.jsとnpmをmiseでインストールする
setup-node:
	@echo "Installing Node.js and npm using mise..."
	@if ! command -v mise &> /dev/null; then \
		echo "Error: mise is not installed. Please install mise first."; \
		exit 1; \
	fi
	@if ! mise plugins list | grep -q node; then \
		echo "Installing node plugin for mise..."; \
		mise plugins install node; \
	fi
	@echo "Installing Node.js LTS version..."
	@mise use --global node@lts
	@echo "Node.js $(mise exec -- node -v) installed"
	@echo "npm $(mise exec -- npm -v) installed"

# TypeScriptクライアントの依存関係をインストールする
setup-ts: setup-node
	@echo "Installing TypeScript dependencies..."
	@mise exec -- npm install

# TypeScriptコードをビルドする
ts-build: setup-ts
	@echo "Building TypeScript client..."
	@mise exec -- npm run build

# TypeScriptサンプルクライアントを実行する
ts-client: ts-build
	@echo "Running TypeScript sample client..."
	@mise exec -- node dist/examples/typescript/client.js

# サーバーを実行
run-server:
	go run cmd/server/main.go

# Go クライアントを実行
run-client:
	go run cmd/client/main.go

# 生成されたファイルをクリーン
clean:
	rm -rf gen/*
	rm -rf dist
	rm -rf node_modules

# 依存関係を更新
deps:
	go mod tidy
	@if [ -f "package.json" ]; then \
		echo "Updating npm dependencies..."; \
		mise exec -- npm update; \
	fi

# テストを実行
test:
	go test ./...
	@if [ -f "package.json" ]; then \
		echo "Running TypeScript tests..."; \
		mise exec -- npm test || echo "No TypeScript tests defined"; \
	fi

# モジュール初期化を実行（初回のみ）
init-module:
	go mod init github.com/takutakahashi/agent-connect

# サンプルリクエストを実行
curl-ping:
	curl \
		--header "Content-Type: application/json" \
		--data '{}' \
		http://localhost:8080/remote.RemoteAgentService/Ping

curl-execute:
	curl \
		--header "Content-Type: application/json" \
		--data '{"session_id":"test-curl","provider":{"model_name":"gpt-4o","api_key":"dummy-key","provider_name":"openai"},"github":{"api_token":"github-token","repo":"takutakahashi/agent-connect","branch_name":"main"},"instruction":"Hello from curl!"}' \
		http://localhost:8080/remote.RemoteAgentService/ExecuteTask

curl-execute-without-github:
	curl \
		--header "Content-Type: application/json" \
		--data '{"session_id":"test-curl","provider":{"model_name":"gpt-4o","api_key":"dummy-key","provider_name":"openai"},"instruction":"Hello from curl without GitHub!"}' \
		http://localhost:8080/remote.RemoteAgentService/ExecuteTask