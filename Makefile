.PHONY: all generate run-server run-client clean

all: generate

# Protoファイルからコードを生成
generate:
	buf generate

# サーバーを実行
run-server:
	go run cmd/server/main.go

# クライアントを実行
run-client:
	go run cmd/client/main.go

# 生成されたファイルをクリーン
clean:
	rm -rf gen/*

# 依存関係を更新
deps:
	go mod tidy

# テストを実行
test:
	go test ./...

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