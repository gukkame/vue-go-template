package userAuth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	db "server/database"
	"golang.org/x/crypto/bcrypt"
)

// Function to create a new user account
func CreateUser(username, password, email, user_type string) error {
    // Generate a random salt
    salt, err := generateRandomSalt(16)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("Password and username: ", username, " ", password)
    // Combine the password and salt
    passwordWithSalt := password + salt

    // Hash the combined value using bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
    if err != nil {
        return err
    } 
    // Store the username, hashed password, and other account details in the database
    stmt, err := db.DBC.Prepare("INSERT INTO Users (username, password, email, salt, user_type, created_at) VALUES (?, ?, ?, ?, ?, ?);")

	  if err != nil {
		  return err
	  }

	  defer stmt.Close() 
	
	  _, err = stmt.Exec(username, hashedPassword, email , salt, user_type, time.Now())
	
	  if err != nil {
		  return err
	  }

    return nil
}

func generateRandomSalt(length int) (string, error) {
    salt := make([]byte, length)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }

    // Encode the salt value to a string representation
    encodedSalt := base64.URLEncoding.EncodeToString(salt)

    return encodedSalt, nil
}
