package structures

type User struct {
	Name       string
	Email      string
	Password   string
	RePassword string
	Error      string
}

type AuthError struct {
}
