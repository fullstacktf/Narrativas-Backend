build:
	docker-compose build
start:
	docker-compose up
stop:
	docker-compose stop
clean:
	rm -rf database/mysql_data
