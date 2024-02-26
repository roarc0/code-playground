all: compose run

.PHONY: run
run:
	go run ./...

.PHONY: compose
compose:
	podman compose -f compose.yaml  up -d


.PHONY: migrate-up
migrate-up:
	migrate -database mysql://root:dbpass@/hackernews -path internal/db/migrations/mysql up

