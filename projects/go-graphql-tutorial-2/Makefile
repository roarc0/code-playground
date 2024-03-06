all: compose run

.PHONY: run
run:
	go run ./...

.PHONY: start
start:
	podman compose -f compose.yaml up -d

.PHONY: stop
stop:
	podman compose -f compose.yaml down

.PHONY: graphql_gen
graphql_gen:
	gqlgen generate
