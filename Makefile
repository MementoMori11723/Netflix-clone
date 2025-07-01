default_goal: run

run:
	@docker-compose -f config/compose.yml up -d

stop:
	@docker-compose -f config/config.yml down --remove-orphans

dev:
	@go run .
