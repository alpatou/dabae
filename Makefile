run:
	go run ./cmd/web -addr=${MYPORT}
	
win-run:
	go run ./cmd/web 

help: 
	go run ./cmd/web -help
	