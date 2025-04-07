include .env
export $(shell sed 's/=.*//' .env)

.PHONY: run format docker-run

run: docker-run
	air

format:
	go fmt ./...

docker-run:
	$(MAKE) -C ../../infra/local docker-run
