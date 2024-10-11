package dbHelper

import (
	"RMS/database"
	"RMS/models"
	"RMS/utils"
	"github.com/jmoiron/sqlx"
	"time"
)

func IsUserExists(email string) (bool, error) {
	SQL := `SELECT count(id) > 0 as is_exist
			  FROM users
			  WHERE email = TRIM($1)
			    AND archived_at IS NULL`

	var check bool
	chkErr := database.RMS.Get(&check, SQL, email)
	return check, chkErr
}

func CreateUser(tx *sqlx.Tx, name, email, password, createdBy string, role models.Role) (string, error) {
	SQL := `INSERT INTO users (name, email, password, created_by, role)
			  VALUES (TRIM($1), TRIM($2), $3, $4, $5) RETURNING id`

	var userID string
	crtErr := tx.Get(&userID, SQL, name, email, password, createdBy, role)
	return userID, crtErr
}

func CreateUserAddress(tx *sqlx.Tx, userID string, addresses []models.AddressRequest) error {
	SQL := `INSERT INTO address (user_id, address, latitude, longitude) VALUES`

	values := make([]interface{}, 0)
	for i := range addresses {
		values = append(values,
			userID,
			addresses[i].Address,
			addresses[i].Latitude,
			addresses[i].Longitude,
		)
	}
	SQL = utils.SetupBindVars(SQL, "(?, ?, ?, ?)", len(addresses))

	_, err := tx.Exec(SQL, values...)
	return err
}

func CreateUserSession(userID string) (string, error) {
	var sessionID string
	SQL := `INSERT INTO user_session(user_id) 
              VALUES ($1) RETURNING id`
	crtErr := database.RMS.Get(&sessionID, SQL, userID)
	return sessionID, crtErr
}

func GetUserInfo(body models.LoginRequest) (string, models.Role, error) {
	SQL := `SELECT u.id,
       			   u.role,
       			   u.password
			  FROM users u
			  WHERE u.email = TRIM($1)
			    AND u.archived_at IS NULL`

	var user models.LoginData
	if getErr := database.RMS.Get(&user, SQL, body.Email); getErr != nil {
		return "", "", getErr
	}
	if passwordErr := utils.CheckPassword(body.Password, user.PasswordHash); passwordErr != nil {
		return "", "", passwordErr
	}
	return user.ID, user.Role, nil
}

func GetArchivedAt(sessionID string) (*time.Time, error) {
	var archivedAt *time.Time

	SQL := `SELECT archived_at 
              FROM user_session 
              WHERE id = $1`

	getErr := database.RMS.Get(&archivedAt, SQL, sessionID)
	return archivedAt, getErr
}

func DeleteUserSession(sessionID string) error {
	SQL := `UPDATE user_session
			  SET archived_at = NOW()
			  WHERE id = $1
			    AND archived_at IS NULL`

	_, delErr := database.RMS.Exec(SQL, sessionID)
	return delErr
}

//func GetAllUsersByAdmin() ([]models.User, error) {
//	query := `
//				SELECT u.id,
//					   u.name,
//					   u.email,
//					   a.address,
//					   ur.role
//				FROM users u
//						 INNER JOIN user_roles ur
//									ON u.id = ur.user_id
//						 INNER JOIN address a
//									ON u.id = a.user_id
//				WHERE u.archived_at IS NULL
//				  AND ur.role = 'user';
//			`
//
//	users := make([]models.User, 0)
//	FetchErr := database.RMS.Select(&users, query)
//	return users, FetchErr
//}
//
//func GetAllUsersBySubAdmin(loggedUserId string) ([]models.User, error) {
//	query := `
//				SELECT u.id,
//					   u.name,
//					   u.email,
//					   a.address,
//					   ur.role
//				FROM users u
//						 INNER JOIN user_roles ur
//									ON u.id = ur.user_id
//						 INNER JOIN address a
//									ON u.id = a.user_id
//				WHERE u.archived_at IS NULL
//				  AND u.created_by = $1
//				  AND ur.role = 'user';
//			`
//
//	users := make([]models.User, 0)
//	FetchErr := database.RMS.Select(&users, query, loggedUserId)
//	return users, FetchErr
//}
