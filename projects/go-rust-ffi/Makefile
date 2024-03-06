ROOT_DIR := $(dir $(realpath $(pwd)))

all: build

.PHONY: build
build: build-dynamic build-static

.PHONY: build-dynamic
build-dynamic: build-libgreet
	go build -o build/dynamic -ldflags="-r $(ROOT_DIR)lib" dynamic.go

.PHONY: build-static
build-static: build-libgreet
	go build -o build/static static.go

.PHONY: build-libgreet
build-libgreet:
	@cd libgreet && cargo build --release
	@mkdir build > /dev/null 2>&1 || true
	@cp libgreet/target/release/libgreet.{so,a} build/

.PHONY: run
run: run-dynamic run-static

.PHONY: run-dynamic
run-dynamic: build-dynamic
	@LD_LIBRARY_PATH=./build:$(LD_LIBRARY_PATH) ./build/dynamic

.PHONY: run-static
run-static: build-static
	@./build/static

.PHONY: clean
clean:
	rm -rf build/ libgreet/target/
