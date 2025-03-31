.PHONY: docker-run docker-stop

docker-run:
	docker compose pull
	docker compose up -d --build --wait

docker-stop:
	docker compose down