run_day01:
	@cat assets/day01/input | go run cmd/day01/main.go
.PHONY: run_day01

test:
	go test ./...
