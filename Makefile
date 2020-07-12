run:
	go run gospec.go --test-files ./test_files

build:
	go build -o gospec.so ./... 