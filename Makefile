.PHONY: lint
# USE DOCKER FOR GENERATING THINGS ON CI

generate_sqlc:
	docker run --rm -v $$(pwd)/internal/storage/postgres:/src -w /src kjconroy/sqlc generate

generate_protos:
	buf generate

lint_protos:
	buf lint

migrate_up:
	migrate -source file://sql/migrations \
	 		-database learnings up 1
