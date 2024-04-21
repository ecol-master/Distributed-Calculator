docker-start:
	docker network create -d bridge custom_netw
	docker compose up -d --build
