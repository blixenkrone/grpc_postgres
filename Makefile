.PHONY: generate_sqlc,generate_protos,lint_protos,migrate_up,postgres_local

generate: generate_protos generate_sqlc

generate_sqlc:
	docker run --rm -v $$(pwd)/storage/postgres:/src -w /src kjconroy/sqlc generate

generate_protos:
	buf generate

lint_protos:
	buf lint

migrate_up:
	migrate -source file://sql/migrations \
	 		-database learnings up 1

postgres_local:
	docker run --rm -d -e POSTGRES_PASSWORD=example -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -e HOST_HOSTNAME=postgres -p 5432:5432 postgres:latest
