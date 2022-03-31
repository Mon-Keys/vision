build:
	rm -rf build
	mkdir build
	go build cmd/vision/main.go
	mv ./main build/


run-docker:
	docker compose up --build

format:
	go fmt ./...

clean-docker:
	docker system prune	-f
	docker volume prune -f
