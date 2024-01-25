migrate:
	/opt/homebrew/bin/migrate -path migrations/ -database "postgres://postgres:1234@127.0.0.1:5432/postgres?sslmode=disable" up

