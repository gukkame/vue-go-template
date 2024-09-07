package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "server/middleware"
	ath "server/services/auth/user-auth"
	db "server/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var PORT string;

// Main function for PostgreSQL 
func main() {
	fmt.Println("Server is running...")

	http.HandleFunc("/login", mw.CORS(ath.HandleLogIn))

	
	// SQLite database
	db.Database()

	// Close database connection when main function gets closed
	defer db.DBC.Close()

	PORT = os.Getenv("PORT")

	if (PORT == "") {
		PORT = "8080"
	}

	directoryPath := "./resources"

    _, err := os.Stat(directoryPath)
    if os.IsNotExist(err) {
        fmt.Printf("Directory '%s' not found.\n", directoryPath)
        return
    }
	// Images -> ./resources
	fileServer := http.FileServer(http.Dir(directoryPath))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))

	// Golang server
	fmt.Printf("API Server running at port "+ PORT +"/\n")
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		fmt.Println("(server.go) Golang server has stopped due to:")
		log.Fatal(err)
	}
}

// Open and create database and register admin user 
func init() {
	// db.Database()

	// registerAdminUsingFlag()	
	// registerAdmin()
}

// Register admin user by env variables
func registerAdmin() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
		os.Exit(0)
    }

	// Access environment variables
	username := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	email := os.Getenv("E-MAIL")

	if err := ath.CreateUser(username, password, email, "admin"); err != nil {
		log.Fatal(err)
	}

}

// Register admin user by terminal
func registerAdminUsingFlag() {
	registerPtr := flag.Bool("register", false, "register a new user with this flag")

	var username string
	flag.StringVar(&username, "username", "bar", "a string var")
	var password string
	flag.StringVar(&password, "password", "bar", "a string var")
	var email string
	flag.StringVar(&email, "email", "bar@bar.lv", "a string var")

	flag.Parse()
	if *registerPtr {
		if err := ath.CreateUser(username, password, email, "admin"); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}