package sessionAuth

type ExistingSession struct {
	Token       string
	Expiry_date string
	User_id     int
}

type Cookie struct {
	Id       string
	Username string
}

type Username struct {
	Username string
	Token    string
}
