SRC_DOCKER_IMAGES := $(shell docker images -q candystore-exercise-src)

up:
	docker-compose up --build --remove-orphans

down:
ifeq ($(strip $(SRC_DOCKER_IMAGES)),)
	@echo "No images to remove"
else
	docker-compose down -v --remove-orphans
	docker rmi $(SRC_DOCKER_IMAGES)
endif

.PHONY: up down
