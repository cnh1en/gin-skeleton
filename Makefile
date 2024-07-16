APP_NAME=server 

dev:
	docker-compose -f docker-compose-dev.yml up
	go run ./cmd/${APP_NAME}