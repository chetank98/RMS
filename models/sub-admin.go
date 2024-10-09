package models

type RegisterSubAdminRequest struct {
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Role      Role   `json:"role" db:"role"`
	CreatedBy string `json:"createdBy" db:"created_by"`
}

type SubAdmin struct {
	Id    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Role  Role   `json:"role" db:"role"`
}
