package user

type NewUser struct {
	FullName    string
	Username    string
	Email       string
	Hash        string
	Preferences string //JSON
}
