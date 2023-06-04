package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"fmt"
)

var Conn *pgxpool.Pool

// DBConnection is What
func PgDBConnection(host, port, user, password, database string) (err error) {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	Conn, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	err = Conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection database: %v\n", err)
		os.Exit(1)
	}

	return nil
}
