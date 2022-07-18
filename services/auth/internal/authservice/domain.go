package authservice

type TokenSession struct {
	Token    string
	Username string
	Email    string
	Role     string
}
