define kill_process
	kill $$(ps | perl -pe 's/^ +//' | grep $(1) | grep -E -e '^[0-9]+' -o)
endef

.PHONY: build
build:
	go build -o mock cmd/mock/*.go
	go build -o app cmd/app/*.go

.PHONY: test
test: build
	./mock &
	go clean -testcache
	go test ./...
	$(call kill_process,mock)
