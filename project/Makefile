BROKER_BINARY=brokerApp

##	up:	starts all containers in the background
up:
	@echo	"Starting docker images..."
	docker-compose up -d
	@echo	"Docker images started!"

##	down:	stops all running containers
down:
	@echo	"Stopping docker images..."
	docker-compose down
	@echo	"Docker stopped!"

##	build_broker:	builds broker binary
build_broker:
	@echo	"Building Broker Binary"
	cd ../broker-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo	"Broker Binary Built!!"

up_build:	build_broker
	@echo	"Stopping Docker Images if running"
	docker-compose down
	@echo	"Building and starting Docker images"
	docker-compose up --build -d
	@echo	"Docker Images built and started!"

