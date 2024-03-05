
migration-up:
	migrate -path ./migrations/postgres -database 'postgres://javohir:12345@0.0.0.0:5432/back?sslmode=disable' up
