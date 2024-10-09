package dbHelper

import (
	"RMS/database"
	"RMS/models"
)

func GetAllSubAdmins() ([]models.SubAdmin, error) {
	query := `
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
	FetchErr := database.RMS.Select(&subAdmins, query)
	return subAdmins, FetchErr
}
