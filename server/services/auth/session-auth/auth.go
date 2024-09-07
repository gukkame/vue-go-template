package sessionAuth

import (
	"net/http"
	db "server/database"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

//ALSO PASTED HERE DUE TO GOLANG RESTRICTON OF NOT ALLOWED TO HAVE PACKAGE IMPORT LOOPS &FROM middleware.go
func SetupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, header1")
}

func DeleteSessionByToken(token string) {
	stmt, _ := db.DBC.Prepare(`DELETE FROM Sessions WHERE token = ?`)
	stmt.Exec(token)
	defer stmt.Close()
}

func GenerateCookieInfo(username, user_type string) Cookie {
	//todo add user type as hashed to cookie info, fe will unhash it 
	newId := uuid.NewV4().String()
	return Cookie{newId, username}
}

func GenerateSession(token_id string, user_id int) error {
	err := DeleteOldSessions(user_id);
	if err != nil {
		return err
	}

	stmt, err := db.DBC.Prepare("INSERT INTO Sessions(token, user_id, expiry_date) VALUES(?, ?, datetime('now','+15 minutes'))")
	if err != nil {
		return err
	}
	stmt.Exec(token_id, user_id)
	defer stmt.Close()
	return err
}

func AuthUser(token string) bool {
	row, err := db.DBC.Query("SELECT * FROM Sessions WHERE token = ?", token)
	if err != nil {
		DeleteSessionByToken(token)
		return false
	}
	defer row.Close()

	var existingSession ExistingSession
	for row.Next() {
		err := row.Scan(
			&existingSession.Token,
			&existingSession.Expiry_date,
			&existingSession.User_id,
		)
		if err != nil {
			DeleteSessionByToken(token)
			return false
		}

	}

	dateNow := time.Now()

	//GETTING THE REQUIRED DATE FORMAT
	dateSession := existingSession.Expiry_date
	cut1 := strings.Split(dateSession, "T")
	join1 := strings.Join(cut1, " ")
	cut2 := strings.Split(join1, "Z")
	join2 := strings.Join(cut2, "")
	DateTimeSession, err := time.Parse("2006-01-02 15:04:05", join2)
	if err != nil {
		DeleteSessionByToken(token)
		return false
	}

	//COMPARING DATETIME NOW TO SESSION EXIPRY_DATE
	if !dateNow.Before(DateTimeSession) {
		DeleteSessionByToken(token)
		return false
	}
	return true
}

func DeleteOldSessions(user_id int) error {
	stmt, err := db.DBC.Prepare(`DELETE FROM Sessions WHERE expiry_date < datetime("now") AND user_id = ?`)
	if err != nil {
		return err
	}

	stmt.Exec(user_id)
	defer stmt.Close()
	return err
}

func CurrentCookieInfo(token string) (string, string) {
	row, _ := db.DBC.Query("SELECT Users.Username, Sessions.token FROM Users INNER JOIN Sessions ON Sessions.user_id = Users.id WHERE token = ?", token)
	defer row.Close()

	var existingUser Username
	for row.Next() {
		err := row.Scan(
			&existingUser.Username,
			&existingUser.Token,
		)
		if err != nil {
			return existingUser.Token, existingUser.Username
		}
	}
	return existingUser.Token, existingUser.Username
}
