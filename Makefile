include .env
export

api-run: 
	go run ./cmd/api/main.go

run:
	go run ./cmd/web -addr=${MYPORT}
	
win-run:
	go run ./cmd/web 

help: 
	go run ./cmd/web -help
	

db-start: 
	docker-compose up -d db

db-stop: 
	docker-compose down 
	
db-connect-root: 
	docker exec -it vasi mysql -uroot -p
	
db-connect-user-web: 
	docker exec -it vasi mysql -D snippetbox -u web -ppass
	
#GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'     IDENTIFIED BY 'YOUR_PASS' WITH GRANT OPTION; FLUSH PRIVILEGES;  

db-create-schema: 
	docker exec -i vasi mysql -uroot -proot  < db/schema.sql
