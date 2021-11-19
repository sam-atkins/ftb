all: install

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

.PHONY: install
install: tidy
	go install

.PHONY: test
test:
	go test ./... -v

.PHONY: cov
cov:
	go test ./... --cover
