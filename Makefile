build_shell:
	docker compose build

shell:
	docker compose run --service-ports --rm uffizzi bash

lint:
	golangci-lint run

fix_lint:
	golangci-lint run --fix
