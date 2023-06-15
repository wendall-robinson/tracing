GOOS=linux
# use the working dir as the app name, this should be the repo name
APP_NAME=$(shell basename $(CURDIR))

.PHONY: build run

docker-up: | build
	@docker-compose build --no-cache --build-arg NAME=${APP_NAME}
	@docker-compose  -f docker-compose.yaml up -d

docker-down:
	@docker-compose -f docker-compose.yaml down

docker-clean:
	@docker-compose -f docker-compose.yaml down --volumes
