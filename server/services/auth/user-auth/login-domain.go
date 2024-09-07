package userAuth

type NewUser struct {
	Username  string
	Email     string
	FirstName string
	LastName  string
	NickName  string
	AboutMe   string
	Age       string
	Gender    string
	Password  string
}

type UserData struct {
	Id           int
	Username     string
	Password     string
	Email        string
	Age          string
	Gender       string
	FirstName    string
	LastName     string
	NickName     string
	AboutMe      string
	Avatar_image string
	User_type    string
	Salt         string
	Created_at   string
}

type Profile struct {
	Username       string
	Email          string
	Age            string
	Gender         string
	FirstName      string
	LastName       string
	NickName       string
	AboutMe        string
	Avatar_image   string
	Date           string
	Status         string
	Profile_status string
}

type LogIn struct {
	Username string
	Password string
	Status   string
}
