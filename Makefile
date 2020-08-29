GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod download

CLIENT=grpcclient
SERVER=grpcserver

client:
	$(GOBUILD) -o $(CLIENT) client/main.go

server:
	$(GOBUILD) -o $(SERVER) server/main.go

clean:
	rm -rf grpcserver grpcclient