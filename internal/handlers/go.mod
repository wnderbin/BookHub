module handlers

go 1.22.2

replace handlers/mysql => ../mysql

require (
	github.com/gorilla/mux v1.8.1
	handlers/mysql v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.0 // indirect
)
