generate:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		todo/todo.proto
get:
	go get ./...
build: get
	echo "Building server..."
	GGO_ENABLED=0 GOOS=linux go build -o bin/linux/server ./server/main.go
	GGO_ENABLED=0 GOOS=darwin go build -o bin/macos/server ./server/main.go
run: build
	echo "Standing up server..."
	docker-compose -p grpc-todoserver -f ./docker-compose.yml up --build