module github.com/gabystallo/go-hexagonal-api-base

go 1.21.13

require github.com/gorilla/mux v1.8.1

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang/mock v1.6.0
	github.com/jmoiron/sqlx v1.4.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

// tool chain:
// go install github.com/golang/mock/mockgen@v1.6.0
// go install gotest.tools/gotestsum@v1.12.0
// go install github.com/vladopajic/go-test-coverage/v2@v2.8.2
// go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6
