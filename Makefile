include .env
export

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
	
db-connect: 
	docker exec -it vasi mysql -uroot -p