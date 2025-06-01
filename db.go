package main

import (
	"context"
	"database/sql"
	"os"

	sql_gen "github.com/ranjbar-dev/tel-bot/sql/gen"
	"github.com/ranjbar-dev/tel-bot/sql/schemas"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func init() {

	// create db file if not exists
	connection, err := sql.Open("sqlite", os.Getenv("DB_PATH"))
	if err != nil {
		panic(err)
	}

	// Check if users table exists
	var tableExists bool
	err = connection.QueryRow("SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type='table' AND name='users')").Scan(&tableExists)
	if err != nil {
		panic(err)
	}

	// Create users table if it doesn't exist
	if !tableExists {
		if _, err := connection.ExecContext(context.Background(), schemas.UsersSchema); err != nil {
			panic(err)
		}
	}

	// set db
	db = connection
}

func findUser(chatID int64) (sql_gen.User, error) {

	queries := sql_gen.New(db)

	return queries.FindUser(context.Background(), chatID)
}

func insertUser(chatID int64, name string) (sql_gen.User, error) {

	queries := sql_gen.New(db)

	return queries.InsertUser(context.Background(), sql_gen.InsertUserParams{
		ChatID: chatID,
		Name:   name,
	})
}

func updateUser(chatID int64, name string) (sql_gen.User, error) {

	queries := sql_gen.New(db)

	return queries.UpdateUserInformation(context.Background(), sql_gen.UpdateUserInformationParams{
		Name:   name,
		ChatID: chatID,
	})
}

func deleteUser(chatID int64) error {

	queries := sql_gen.New(db)

	return queries.DeleteUser(context.Background(), chatID)
}
