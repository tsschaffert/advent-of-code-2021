run_day01:
	@cat assets/day01/input | go run cmd/day01/main.go
.PHONY: run_day01

run_day02:
	@cat assets/day02/input | go run cmd/day02/main.go
.PHONY: run_day02

test:
	go test ./...
