dev:
	cd cmd\http_api\ && go run .

remove_mysql:
	docker rm -f go_todo_mysql
	
create_mysql:
	docker run --name go_todo_mysql -p 3457:3306 -e MYSQL_ROOT_PASSWORD=mypassword -d mysql:8
	
