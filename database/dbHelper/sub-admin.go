package dbHelper

import (
	"RMS/database"
	"RMS/models"
)

func CreateSubAdmin(name, email, password, createdBy string, role models.Role) error {
	SQL := `INSERT INTO users (name, email, password, created_by, role)
			  VALUES (TRIM($1), TRIM($2), $3, $4, $5) RETURNING id`

	var userID string
	crtErr := database.RMS.Get(&userID, SQL, name, email, password, createdBy, role)
	return crtErr
}

func GetAllSubAdmins() ([]models.SubAdmin, error) {
	SQL := `SELECT id,
				   name,
				   email,
				   role,
				   created_by
			FROM users
				WHERE role = 'sub-admin' 
				AND archived_at IS NULL`

	subAdmins := make([]models.SubAdmin, 0)
	fetchErr := database.RMS.Select(&subAdmins, SQL)
	return subAdmins, fetchErr
}
