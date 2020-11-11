run:
	go run ./... -file test_files/anyhost.yaml

test:
	go test ./... -v