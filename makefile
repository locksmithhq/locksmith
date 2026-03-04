.PHONY: help up down restart logs logs-api logs-db clean build rebuild

# Variáveis
ENV_FILE=.env

# Cores para output
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

include $(ENV_FILE)

help: ## Mostra esta mensagem de ajuda
	@echo "$(GREEN)Locksmith - Comandos disponíveis:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

up: ## Inicia todos os serviços (database + api) na mesma network
	@echo "$(GREEN)Iniciando serviços...$(NC)"
	@docker compose -f compose.yaml --env-file $(ENV_FILE) up --build -d
	@echo "$(GREEN)Todos os serviços estão rodando!$(NC)"
	@echo "$(YELLOW)Use 'make logs' para ver os logs$(NC)"
	@echo "$(YELLOW)Use 'make logs-db' para ver os logs do banco de dados$(NC)"
	@echo "$(YELLOW)Use 'make logs-pgweb' para ver os logs do pgweb$(NC)"
	@echo "$(YELLOW)Use 'make logs-api' para ver os logs da API$(NC)"
	@echo "$(YELLOW)Use 'make logs-web' para ver os logs da web$(NC)"
	@echo "$(YELLOW)Use 'make logs-docs' para ver os logs da documentação$(NC)"
	@echo "$(YELLOW)Use 'make open-pgweb' para abrir o pgweb no navegador$(NC)"
	@echo "$(YELLOW)Use 'make open-web' para abrir a web no navegador$(NC)"
	@echo "$(YELLOW)Use 'make open-api' para abrir a API no navegador$(NC)"
	@echo "$(YELLOW)Use 'make open-docs' para abrir a documentação no navegador$(NC)"

down: ## Para todos os serviços
	@docker compose -f compose.yaml --env-file $(ENV_FILE) down
	@echo "$(GREEN)Todos os serviços foram parados$(NC)"

restart: down up ## Reinicia todos os serviços

logs: ## Mostra logs de todos os serviços
	@docker compose -f compose.yaml --env-file $(ENV_FILE) logs -f

logs-api: ## Mostra logs apenas da API
	@docker compose -f compose.yaml --env-file $(ENV_FILE) logs api -f

logs-web: ## Mostra logs apenas da web
	@docker compose -f compose.yaml --env-file $(ENV_FILE) logs web -f

logs-db: ## Mostra logs apenas do banco de dados
	@docker compose -f compose.yaml --env-file $(ENV_FILE) logs database -f

clean: down ## Para os serviços e remove volumes
	@echo "$(YELLOW)Removendo volumes...$(NC)"
	@docker compose -f compose.yaml --env-file $(ENV_FILE) down -v
	@echo "$(GREEN)Limpeza concluída$(NC)"

build: ## Reconstrói as imagens Docker
	@echo "$(GREEN)Reconstruindo imagem da API...$(NC)"
	@docker compose -f compose.yaml --env-file $(ENV_FILE) build

rebuild: down build up ## Para, reconstrói e inicia todos os serviços

status: ## Mostra o status dos containers
	@echo "$(GREEN)Status dos containers:$(NC)"
	@docker ps -a --filter "name=database" --filter "name=api" --filter "name=web" --filter "name=proxy" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

shell-api: ## Abre shell no container da API
	@docker exec -it api /bin/sh

shell-db: ## Abre shell no container do banco de dados
	@docker exec -it database psql -U $${POSTGRES_USER:-postgres} -d $${POSTGRES_DB:-locksmith}

prune: ## Remove containers, networks e imagens não utilizadas
	@echo "$(YELLOW)Removendo recursos Docker não utilizados...$(NC)"
	@docker system prune -f
	@echo "$(GREEN)Limpeza do sistema concluída$(NC)"

open-pgweb: ## Abre o pgweb no navegador
	@open http://localhost:10002

open-web: ## Abre a web no navegador
	@open http://localhost:10000

open-api: ## Abre a API no navegador
	@open http://localhost:10001

open-docs: ## Abre a documentação no navegador
	@open http://localhost:${DOCS_PORT}

build-prod:
	@echo "$(GREEN)Preparando build de produção...$(NC)"
	@cp -r api locksmith/
	@cp -r config locksmith/
	@cp -r web locksmith/
	@echo "$(GREEN)Arquivos copiados para a pasta locksmith$(NC)"
	@docker buildx build --platform linux/amd64,linux/arm64 -t booscaaa/locksmith:latest -f locksmith/Dockerfile . --push
	@rm -rf locksmith/api locksmith/config locksmith/web
	@echo "$(GREEN)Arquivos temporários removidos da pasta locksmith$(NC)"
	



	
