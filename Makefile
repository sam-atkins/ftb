all: install

.PHONY: install
install:
	go mod vendor -v
	go install

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor
