test: 
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

run:
	go run main.go

docker:
	docker build -t tictactoe .

run-docker: docker
	docker run -p 3000:3000 tictactoe