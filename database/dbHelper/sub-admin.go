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
	SQL := `
				SELECT u.id,
					   u.name,
					   u.email,
					   ur.role
				FROM users u
						 INNER JOIN user_roles ur
									ON u.id = ur.user_id
				WHERE u.archived_at IS NULL
				  AND ur.role = 'sub-admin';
			`

	subAdmins := make([]models.SubAdmin, 0)
	FetchErr := database.RMS.Select(&subAdmins, SQL)
	return subAdmins, FetchErr
}
