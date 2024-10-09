package models

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleSubAdmin Role = "sub-admin"
	RoleUser     Role = "user"
)

func (r Role) IsValid() bool {
	return r == RoleAdmin || r == RoleSubAdmin || r == RoleUser
}

type RegisterUserRequest struct {
	Name     string           `json:"name" validate:"required"`
	Email    string           `json:"email" validate:"email"`
	Password string           `json:"password" validate:"gte=6,lte=15"`
	Role     []Role           `json:"role" validate:"required"`
	Address  []AddressRequest `json:"address" validate:"required"`
}

type AddressRequest struct {
	Address   string `json:"address" validate:"required"`
	Latitude  string `json:"latitude" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"gte=6,lte=15"`
}

type Address struct {
	ID        string `json:"id" validate:"id"`
	Address   string `json:"address" db:"address"`
	Latitude  string `json:"latitude" db:"latitude"`
	Longitude string `json:"longitude" db:"longitude"`
}

type User struct {
	ID      string    `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Email   string    `json:"email" db:"email"`
	Address []Address `json:"address" db:"address"`
	Role    []Role    `json:"role" db:"role"`
}

type UserCtx struct {
	UserID    string `json:"userId"`
	SessionID string `json:"sessionId"`
	Role      []Role `json:"role"`
}
