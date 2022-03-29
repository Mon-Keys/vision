build:
	rm -rf build
	mkdir build
	go build cmd/proxy/main.go
	mv ./main build/


run-docker:
	docker-compose up


clean-docker:
	docker system prune	-f
	docker volume prune -f
