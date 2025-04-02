include .env
export $(shell sed 's/=.*//' .env)

.PHONY: run docker-run

run: docker-run
	go run main.go


docker-run:
	$(MAKE) -C ../../infra/local docker-run