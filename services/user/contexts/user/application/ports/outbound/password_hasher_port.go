package outbound

type PasswordHasherPort interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) bool
}
