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
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Address   string `json:"address" db:"address"`
	Role      Role   `json:"role" db:"role"`
	CreatedBy string `json:"createdBy" db:"created_by"`
}

type LoginUser struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type User struct {
	Id      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Email   string `json:"email" db:"email"`
	Address string `json:"address" db:"address"`
	Role    Role   `json:"role" db:"role"`
}

type UserCtx struct {
	UserId    string `json:"userId" db:"user_id"`
	SessionId string `json:"sessionId" db:"session_id"`
	Role      Role   `json:"role" db:"role"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
}

type UserInfo struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Role     Role   `json:"role" db:"role"`
}
