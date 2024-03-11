DOCKER := docker

all: compose run

.PHONY: run
run:
	go run ./...

.PHONY: start
start:
	${DOCKER} compose up -d

.PHONY: stop
stop:
	${DOCKER} compose down