build:
	cd cmd && go build -o ../bin/main

run:
	build && bin/main

run-dev:
	go run cmd/main.go