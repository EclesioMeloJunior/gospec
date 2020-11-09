run:
	go run ./... -test test_files/anyhost.yaml

test:
	go test ./... -v