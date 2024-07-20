package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func hello() {
	dbUrl := "postgres://Golang1_admin:12345@localhost:5432/golang1"
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connection successful")

}
