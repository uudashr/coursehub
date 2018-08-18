package account

// AccountCreated is events raised when new account created.
type AccountCreated struct {
	ID    string
	Name  string
	Email string
}
