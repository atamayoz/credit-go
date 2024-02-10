package db

import (
	"entgo.io/ent/dialect"
	"github.com/atamayoz/credit-go/ent"
	_ "github.com/lib/pq"
)

func Connect() (*ent.Client, error) {
	//TODO: Improve the connection to receive params
	client, err := ent.Open(dialect.Postgres, "postgres://user:admin@localhost:54320/credit_golang?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return client, err
}
