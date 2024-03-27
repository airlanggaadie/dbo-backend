package configuration

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func (c *configuration) initPostgreSql() *configuration {
	fmt.Println("setting up database connection...")
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		dbURL = "postgres://postgres:postgres@localhost:5432/dbo?sslmode=disable"
	}

	// setup the database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("[configuration][initPostgreSql] failed to connect to database: %v", err)
	}

	c.DB = db

	return c
}

func (c *configuration) migrate() *configuration {
	fmt.Println("running migration...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := c.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err != nil {
		log.Fatalf("[configuration][migrate] begin transaction error: %v", err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS "users" (
			id                  UUID PRIMARY KEY NOT NULL,
			username            VARCHAR(20) NOT NULL UNIQUE,
			name				VARCHAR(50) NOT NULL,
			created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,
		`COMMENT ON COLUMN "users".username IS 'username for authentication';`,
		`CREATE TABLE IF NOT EXISTS "user_password" (
			user_id             UUID NOT NULL UNIQUE,
			password			VARCHAR(50) NOT NULL,
			created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES "users" (id)
		);`,
		`CREATE TABLE IF NOT EXISTS "user_login_history" (
			user_id             UUID NOT NULL,
			created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES "users" (id)
		);`,
		`CREATE TABLE IF NOT EXISTS "orders" (
			id             		UUID PRIMARY KEY NOT NULL,
			code				VARCHAR(20) NOT NULL,
			buyer_name			VARCHAR(50) NOT NULL,
			item_quantity 		BIGINT NOT NULL,
			total_price			DECIMAL NOT NULL,
			created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at			TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,
		`COMMENT ON COLUMN "orders".code IS 'order number based on operational';`,
		`COMMENT ON COLUMN "orders".item_quantity IS 'total item quantity in one order';`,
		`COMMENT ON COLUMN "orders".total_price IS 'total price for order';`,
		`CREATE TABLE IF NOT EXISTS "order_item" (
			id             		UUID PRIMARY KEY NOT NULL,
			order_id			UUID NOT NULL,
			code				VARCHAR(20) NOT NULL,
			name				VARCHAR(50) NOT NULL,
			quantity			BIGINT NOT NULL,
			unit_price			BIGINT NOT NULL,
			total_price			BIGINT NOT NULL,
			created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at			TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (order_id) REFERENCES "orders" (id)
		);`,
		`COMMENT ON COLUMN "order_item".code IS 'item code';`,
		`COMMENT ON COLUMN "order_item".unit_price IS 'price per unit';`,
		`COMMENT ON COLUMN "order_item".total_price IS 'total price means unit price times quantity';`,
	}

	for i, query := range queries {
		_, err = tx.ExecContext(
			ctx,
			query,
		)
		if err != nil {
			tx.Rollback()
			log.Fatalf("[configuration][migrate] execution [%d] error: %v", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatalf("[configuration][migrate] commit error: %v", err)
	}

	return c
}
