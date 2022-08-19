
build:
	docker-compose build

run: build
	docker-compose up

down:
	docker-compose down

re: clean build
	docker-compose up


clean:
	rm -rf ./deployment/data
