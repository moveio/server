
all:
	docker-compose build

server:
	docker-compose up server

dev:
	docker-compose up devserver

test:
	docker-compose up test