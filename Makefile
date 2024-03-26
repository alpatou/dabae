run:
	go run ./cmd/web -addr=":9999"
	
dry-run:
	go run ./cmd/web 

help: 
	go run ./cmd/web -help
	