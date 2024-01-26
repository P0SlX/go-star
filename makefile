MAIN_FILE=main.go
OUTPUT_FILE=application

run:
	go run $(MAIN_FILE) -img=$(img)

run-build: build
	./dist/$(OUTPUT_FILE) -img=$(img)

clean-cache:
	go clean -cache -testcache -modcache

build: clean-cache
	go build -o dist/$(OUTPUT_FILE) -x

time: build
	time ./dist/$(OUTPUT_FILE)

benchmark:
	go test -bench=. benchmarks