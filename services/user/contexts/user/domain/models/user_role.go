package models

type Role int

const (
	RoleAdmin Role = iota
	RoleUser
	RoleGuest
)

func NewRole(role int) Role {
	return Role(role)
}

func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleUser:
		return "user"
	case RoleGuest:
		return "guest"
	default:
		return "unknown"
	}
}

func (r Role) IsValid() bool {
	return r >= RoleAdmin && r <= RoleGuest
}
