package user

type User struct {
	Username string
	Email    string
	Password string
	SessionID string
}

func NewUser(username, email, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}
func (u *User) SetSessionID(sessionID string) {
	u.SessionID = sessionID
}