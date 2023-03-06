pb:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative part1/proto/*.proto

run-grpc:
	@go run main.go grpc development

run-consumer:
	@go run main.go consumer development

redis-cli:
	@docker run --network=container:redis -it redis redis-cli -h redis -a testing123

producer:
	@docker exec -it redpanda-0 rpk topic produce process_ohlc

start:
	@docker-compose up -d
	@python -m webbrowser "http://localhost:6969"