run:
	go run ./cmd/web -addr=$PORT
	
dry-run:
	go run ./cmd/web 

help: 
	go run ./cmd/web -help
	