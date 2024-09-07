package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	migrate "github.com/rubenv/sql-migrate"
)

//Main Database connection for global use
var DBC *sql.DB

func Database() {
	// DBC = openDB()
	// createUserTable(DBC)	// create user database table (if doesnt exist) to add admin user
	// migrations(DBC)

	// SQLite database and migrations
	createDatabaseSQLite()
	DBC = openSQLiteDB()
	migrationsSQLite(DBC) 

	fmt.Println("Connection to database establised")
}

func createDatabaseSQLite() {
	file, err := os.Create("./database/database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Println("database.db created")
}

func migrationsSQLite(db *sql.DB) {

	//TAKES `../database/migrations` DIR AS A SOURCE OF MIGRATIONS
	migrations := &migrate.FileMigrationSource{
		Dir: "./database/migrations",
	}

	// Execute migrations (db conn pool, name, source, action)
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error applying migrations: db.go")
	}
	fmt.Printf("Applied %d migrations to database.db!", n)
	fmt.Println()
}

// Opend SQLite local database
func openSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		fmt.Println("(db.go) Unable to open db due to:")
		log.Fatal(err)
	}

    // Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("(db.go) Unable to ping due to:")
		log.Fatal(err)
	}

	return db
}


func migrations(db *sql.DB) {

	//TAKES `../database/migrations` DIR AS A SOURCE OF MIGRATIONS
	migrations := &migrate.FileMigrationSource{
		Dir: "./database/migrations",
	}

	//EXECUTES THE MIGRATION (db conn pool, name, source, action)
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error applying migrations")	
	}

	fmt.Printf("Applied %d migrations to database.db!", n)
	fmt.Println()
}

// Opening Postgresql database connection
 func openDB() *sql.DB {
    connStr := "user=doadmin password={password} dbname=defaultdb host={host} port=25060 sslmode=require"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database!")

	return db
}

func createUserTable(db *sql.DB)  {
	// Create the Users table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(255) NOT NULL UNIQUE,
            email VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            salt TEXT NOT NULL,
            user_type TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL
        );
    `)
    if err != nil {
        log.Fatal(err)
    }
}