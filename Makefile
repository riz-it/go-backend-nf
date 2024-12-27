create_table:
	migrate -database "postgresql://postgres:postgres@localhost:5432/nurul_faizah?sslmode=disable" -path internal/database/migration up
drop_table:
	migrate -database "postgresql://postgres:postgres@localhost:5432/nurul_faizah?sslmode=disable" -path internal/database/migration down
