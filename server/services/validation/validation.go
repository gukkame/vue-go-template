package validation

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ValidUsername(username string) bool {
	re := regexp.MustCompile(`^([a-zA-Z]+([0-9]*)){4,20}$`)
	return re.MatchString(username)
}

func ValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

//FIRSTNAME & LASTNAME
func ValidNames(name string) bool {
	re := regexp.MustCompile(`^([a-zA-Z\s]){1,20}$`)
	return re.MatchString(name)
}

func ValidAge(age string) bool {
	re := regexp.MustCompile(`\d`)
	if !re.MatchString(age) {
		return false
	}

	intAge, err := strconv.Atoi(age)
	if err != nil {
		return false
	}

	if intAge < 18 || intAge > 100 {
		return false
	}
	return true
}

func ValidGender(gender string) bool {
	re := regexp.MustCompile("^(Male|Female|Other|Unspecified)$")
	return re.Match([]byte(gender))
}

func ValidPassword(password string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	if len(password) >= 8 && len(password) <= 24 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}

func ValidPostTitle(title string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9\s\.,;:!?()"'%\-]{1,60}$`)
	return re.Match([]byte(title))
}

func ValidPostDescription(description string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9\s\.,;:!?()"'%\-]{1,1000}$`)
	return re.Match([]byte(description))
}

func ValidComment(comment string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9\s\.,;:!?()"'%\-]{1,600}$`)
	return re.Match([]byte(comment))
}

//Optional
func ValidNickName(nickname string) bool {
	if nickname == "" {
		return true
	}
	re := regexp.MustCompile(`^[A-Za-z0-9\s\.,;:!?()"'%\-]{1,16}$`)
	return re.Match([]byte(nickname)) 
}
func ValidAboutMe(content string) bool {
	if content == "" {
		return true
	}
	re := regexp.MustCompile(`^[A-Za-z0-9\s\.,;:!?()"'%\-]{1,500}$`)
	return re.Match([]byte(content))
}

func ValidateUserData(username string, email string, firstname string, lastname string, age string, gender string, password string, nickname string, about_me string) bool {
	if !ValidUsername(username) {
		return false
	}

	if !ValidEmail(email) {
		return false
	}

	if !ValidNames(firstname) {
		return false
	}

	if !ValidNames(lastname) {
		return false
	}

	if !ValidAge(age) {
		return false
	}

	if !ValidGender(gender) {
		return false
	}

	if !ValidPassword(password) {
		return false
	}
	if !ValidNickName(nickname) {
		return false
	}
	if !ValidAboutMe(about_me) {
		return false
	}
	return true
}

func ValidateLoginData(username string, password string) bool {

	//CHECKING IF ITS EMAIL OR NOT
	if strings.Contains(username, "@") {
		if !ValidEmail(username) {
			return false
		}
	} else {
		if !ValidUsername(username) {
			return false
		}
	}

	// if !ValidPassword(password) {
	// 	return false
	// }
	return true
}

func ValidatePostData(title string, description string) bool {
	if !ValidPostTitle(title) {
		return false
	}

	if !ValidPostDescription(description) {
		return false
	}
	return true
}