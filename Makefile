	@echo "============= starting db locally ============="
	go mod tidy
	docker-compose -f resources/docker/database/docker-compose.yaml up lopmartyn_container