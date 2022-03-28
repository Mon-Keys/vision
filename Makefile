build:
	rm -rf build
	mkdir build
	go build cmd/proxy/main.go
	mv ./main build/


run:
	go run ./...
