all: install

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

.PHONY: install
install: tidy
	go install
