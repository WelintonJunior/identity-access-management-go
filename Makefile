build:
	@make swag
	@go build -o ./bin/api_server
	@echo "Build finalizado."

run:
	@make build
	@./run.sh

sql_up:
	@echo "Removendo container anterior (se existir)..."
	-@docker rm -f postgres

	@echo "Iniciando novo container postgres com volume persistente..."
	@docker run --name postgres \
		-e POSTGRES_PASSWORD=1234 \
		-e POSTGRES_DB=iam \
		-v postgres_data:/var/lib/postgresql/data \
		-p 5432:5432 \
		-d postgres

	@echo "Aguardando inicialização do container postgres..."
	@sleep 10


redis_up:
	@echo "Iniciando container redis..."
	@docker run --name redis -p 6379:6379 -d redis redis-server --requirepass "" --appendonly yes
	@echo "Aguardando inicialização do container redis..."
	@sleep 10

kafka_up:
	@echo "Iniciando container kafka..."
	@docker run -d --name=kafka -p 9092:9092 apache/kafka
	@echo "Aguardando inicialização do container kafka..."
	@sleep 10

migrate_up:
	@make build
	@go run main.go migrate -o up

migrate_down:
	@make build
	@go run main.go migrate -o down

swag:
	@echo "Atualizando documentação"
	@swag init
	@echo "Documentação atualizada"


