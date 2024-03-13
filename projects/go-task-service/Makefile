BUILD_DIR := build

.PHONY: all
all: test

.PHONY: start
start:
	go run ./...

.PHONY: compose_up
compose_up:
	podman compose -f docker-compose.yaml up -d

.PHONY: compose_down
compose_down:
	podman compose -f docker-compose.yaml down

.PHONY: update
update:
	@go get -u -v ./...
	@go mod tidy

COV_DIR := $(BUILD_DIR)/coverage
COV_FILE := $(COV_DIR)/coverage.out
COV_PKGS := $(shell go list ./... | egrep -v '(/cmd/proto$$)' | tr '\n' ',')
.PHONY: test
test:
	@echo "> test setup"
	@go generate ./...
	@find . -type d -name '*gomock_reflect_*' -exec rm -rf \;
	@mkdir -p ${COV_DIR}
	@echo "> go unit tests"
	@GOEXPERIMENT=nocoverageredesign go test -covermode=count -coverprofile=${COV_FILE}.tmp -v -count=1 -coverpkg=${COV_PKGS} ./...
	@cat ${COV_FILE}.tmp | egrep -v '(_mock.go)' > ${COV_FILE}
	@go tool cover -html=${COV_FILE} -o ${COV_DIR}/index.html
	@go tool cover -func=${COV_FILE} | grep "total:"

.PHONY: lint
lint:
	@echo '> go vet'
	@go vet ./...
	@echo '> golint'
	@golint -set_exit_status ./...

.PHONY: race
race:
	@go test -race -short ./...

.PHONY: msan
msan:
	@go test -msan -short ./...

.PHONY: build_image
build_image:
	docker buildx build -f Dockerfile -t app .

PROF_PACKAGE := ...

.PHONY: profile
profile:
	@go test -o ${BUILD_DIR}/app.test -memprofile ${BUILD_DIR}/memprofile.out -cpuprofile ${BUILD_DIR}/profile.out ${PROF_PACKAGE}
	@go tool pprof -http=":8000" ${BUILD_DIR}/app.test ${BUILD_DIR}/profile.out

.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

