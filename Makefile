export GOLANG_IMAGE := golang:1.12.5-alpine3.9

# Build/ForceRebuild the services
build:
	@docker-compose build --no-cache
.PHONY: build

# Launches the nit tests
unit-tests :
	@docker-compose run --rm $@ || docker-compose logs
.PHONY: unit-tests

# Launches the integration tests
integration-tests:
	@docker-compose up --force-recreate --always-recreate-deps \
		--abort-on-container-exit --exit-code-from $@ $@ \
		|| docker-compose logs
.PHONY: integration-tests