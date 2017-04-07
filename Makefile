
all:
	docker-compose build

server:
	docker-compose up server

devserver:
	docker-compose up devserver