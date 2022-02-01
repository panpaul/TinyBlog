package model

type Role int

const (
	RoleGuest Role = iota
	RoleUser
	RoleAdmin
)

func (r Role) String() string {
	strings := [...]string{"Guest", "User", "Admin"}
	if r < RoleGuest || r > RoleAdmin {
		return "Unknown"
	}
	return strings[r]
}

type CasbinRule struct {
	Role   Role
	Path   string
	Method string
}
