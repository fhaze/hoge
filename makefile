define kill_process
	kill $$(ps | perl -pe 's/^ +//' | grep $(1) | grep -E -e '^[0-9]+' -o)
endef

.PHONY: setup
setup:
	go version
	go mod download

.PHONY: build
build:
	go build -o mock cmd/mock/main.go
	go build -o app cmd/app/main.go

.PHONY: test
test: build
	./mock &
	go clean -testcache
	go test ./...
	$(call kill_process,mock)
