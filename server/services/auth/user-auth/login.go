package userAuth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	db "server/database"
	mw "server/middleware"
	ath "server/services/auth/session-auth"
	val "server/services/validation"
)

var logInData LogIn

func HandleLogIn(w http.ResponseWriter, r *http.Request) {
	mw.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(reqBody), &logInData)

		fmt.Println(logInData)

		if val.ValidateLoginData(logInData.Username, logInData.Password) {

			var row *sql.Rows
			if strings.Contains(logInData.Username, "@") {
				row, err = db.DBC.Query("SELECT * FROM Users WHERE email = ?", logInData.Username)
				if err != nil {
					w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
					return
				}

			} else {
				row, err = db.DBC.Query("SELECT * FROM Users WHERE username = ?", logInData.Username)
				if err != nil {
					w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
					return
				}
			}
			defer row.Close()

			var dbUserData UserData
			for row.Next() {
				err := row.Scan(
					&dbUserData.Id,
					&dbUserData.Username,
					&dbUserData.Email,
					&dbUserData.Password,
					&dbUserData.Salt,
					&dbUserData.User_type,
					&dbUserData.Created_at,
				)
				if err != nil {
					w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
					return
				}
			}

			if dbUserData.Id == 0 {
				w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
				return
			}

			err = row.Err()
			if err != nil {
				w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
				return
			}

			if !mw.CheckPasswordHash(logInData.Password, dbUserData.Password, dbUserData.Salt) {
				w.Write([]byte(`{"error":"LogInError", "message": "Invalid password"}`))
				return
			}

			newCookie := ath.GenerateCookieInfo(dbUserData.Username, dbUserData.User_type)
			error := ath.GenerateSession(newCookie.Id, dbUserData.Id)
			if error != nil {
				w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
				return
			}

			var cookieData []byte
			cookieData, _ = json.Marshal(newCookie)
			w.Write(cookieData)

		} else {
			w.Write([]byte(`{"error":"LogInError", "message": "Invalid credentials"}`))
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"LogInError", "message": "Can't find method requested"}`))
	}
}