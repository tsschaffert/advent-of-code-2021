run_day01:
	@cat assets/day01/input | go run cmd/day01/main.go
.PHONY: run_day01

run_day02:
	@cat assets/day02/input | go run cmd/day02/main.go
.PHONY: run_day02

run_day03:
	@cat assets/day03/input | go run cmd/day03/main.go
.PHONY: run_day03

run_day06:
	@cat assets/day06/input | go run cmd/day06/main.go
.PHONY: run_day06

run_day07:
	@cat assets/day07/input | go run cmd/day07/main.go
.PHONY: run_day07

test:
	go test ./...
