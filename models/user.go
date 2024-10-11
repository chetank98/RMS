package models

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleSubAdmin Role = "sub-admin"
	RoleUser     Role = "user"
)

type UserRequest struct {
	Name     string           `json:"name" validate:"required"`
	Email    string           `json:"email" validate:"email"`
	Password string           `json:"password" validate:"gte=6,lte=15"`
	Address  []AddressRequest `json:"address" validate:"required"`
}

type AddressRequest struct {
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"gte=6,lte=15"`
}

type LoginData struct {
	ID           string `db:"id"`
	PasswordHash string `db:"password"`
	Role         Role   `db:"role"`
}

type Address struct {
	ID        string  `json:"id" db:"id"`
	Address   string  `json:"address" db:"address"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

type User struct {
	ID      string    `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Email   string    `json:"email" db:"email"`
	Address []Address `json:"address" db:"address"`
	Role    Role      `json:"role" db:"role"`
}

type UserCtx struct {
	UserID    string `json:"userId"`
	SessionID string `json:"sessionId"`
	Role      Role   `json:"role"`
}
