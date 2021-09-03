all: build

.PHONY: build
build:
	go mod vendor -v
	go install

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

.PHONY: install
install:
	go install
