DOCKER := docker

all: run

.PHONY: run
run:
	go run ./...

.PHONY: build
build:
	go build ./... ./cmd

.PHONY: start
start:
	${DOCKER} compose up -d

.PHONY: stop
stop:
	${DOCKER} compose down

.PHONY: build_image
build_image:
	${DOCKER} buildx build -f Dockerfile -t roarc/gotmpl-app .
	${DOCKER} push roarc/gotmpl-app