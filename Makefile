build:
	docker-compose build
start:
	docker-compose up
stop:
	docker-compose stop
clean:
	sudo rm -rf database/mysql_data
