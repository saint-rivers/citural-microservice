all:
	docker compose \
		-f keycloak-service/docker-compose.yml \
		-f tinker-service/docker-compose.yml \
		-f rift/docker-compose.yml \
		up -d --build
