.PHONY: up
up:
	docker-compose up --build --always-recreate-deps

api-dev:
	cd ./apps/api && go run ./cmd

ui-dev:
	cd ./apps/web-ui && npm run dev
