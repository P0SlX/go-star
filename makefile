MAIN_FILE=main.go
OUTPUT_FILE=application

run:
	go run $(MAIN_FILE)

run-build:
	./dist/$(OUTPUT_FILE)

clean-cache:
	go clean -cache -testcache -modcache

build: clean-cache
	go build -o dist/$(OUTPUT_FILE) -x