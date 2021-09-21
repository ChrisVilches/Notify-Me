.PHONY: build

build:
	go build -o build/notify_me cli.go notifier.go

run:
	go run cli.go notifier.go
