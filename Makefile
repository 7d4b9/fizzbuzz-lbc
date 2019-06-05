export GOLANG_IMAGE := golang:1.12.5-alpine3.9

TESTS_TARGETS := unit-tests integration-tests

# Launches the  tests
$(TESTS_TARGETS):
	docker-compose run --rm $@ || docker-compose logs
.PHONY: $(TESTS_TARGETS)