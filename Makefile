docker-build:
	docker build -t signature-server .
	
docker-run:
	docker run --publish 3001:3001 --name signature-server --rm signature-server
	
go-run:
	go run main.go

keygen:
	go run utils/generate.go