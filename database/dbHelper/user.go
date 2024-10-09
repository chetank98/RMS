package dbHelper

import (
	"RMS/database"
	"RMS/models"
	"RMS/utils"
	"github.com/jmoiron/sqlx"
	"time"
)

func IsUserExists(email string) (bool, error) {
	query := `
				SELECT count(id) > 0 as is_exist
				FROM users
				WHERE email = TRIM($1)
				  AND archived_at IS NULL;
			`
	var check bool
	chkErr := database.RMS.Get(&check, query, email)
	if chkErr != nil {
		return false, chkErr // Return error if the query fails
	}
	return check, nil
}

func GetUserInfo(email, password string) (*models.UserInfo, error) {
	query := `
				SELECT u.id,
					   u.name,
					   u.password,
					   ur.role
				FROM users u
						 INNER JOIN user_roles ur
									ON u.id = ur.user_id
				WHERE u.archived_at IS NULL
				  AND u.email = TRIM($1);
			`
	var userInfo models.UserInfo
	getErr := database.RMS.Get(&userInfo, query, email)

	if getErr != nil {
		return nil, getErr
	}

	if passwordErr := utils.CheckPassword(password, userInfo.Password); passwordErr != nil {
		return nil, passwordErr
	}
	return &userInfo, nil
}

func CreateUserSession(userId string) (string, error) {
	query := `
				INSERT INTO user_session(user_id)
				VALUES ($1)
				RETURNING id;
			`
	var sessionId string
	crtErr := database.RMS.QueryRowx(query, userId).Scan(&sessionId)
	if crtErr != nil {
		return "", crtErr // Return error if the query fails
	}
	return sessionId, nil
}

func CreateUser(db sqlx.Ext, name, email, password, createdBy string) (string, error) {
	query := `
				INSERT INTO users (name, email, password, created_by)
				VALUES (TRIM($1), TRIM($2), $3, $4)
				RETURNING id;
			`
	var userId string
	err := db.QueryRowx(query, name, email, password, createdBy).Scan(&userId)
	if err != nil {
		return "", err // Return error if the query fails
	}
	return userId, nil
}

func CreateUserRole(db sqlx.Ext, userId string, role models.Role) error {
	query := `
				INSERT INTO user_roles(user_id, role)
				VALUES ($1, $2);
			`
	_, err := db.Exec(query, userId, role)
	return err
}

func CreateUserAddress(db sqlx.Ext, userId, address string) error {
	query := `
				INSERT INTO address(user_id, address)
				VALUES ($1, TRIM($2));
			`
	_, err := db.Exec(query, userId, address)
	return err
}

func GetArchivedAt(sessionId string) (*time.Time, error) {
	query := `
				SELECT archived_at
				FROM user_session
				WHERE id = $1
				  AND archived_at IS NULL;
			`
	var archivedAt *time.Time
	getErr := database.RMS.Get(&archivedAt, query, sessionId)
	if getErr != nil {
		return nil, getErr // Return error if the query fails
	}
	return archivedAt, nil
}

func DeleteUserSession(sessionId string) error {
	query := `
				UPDATE user_session
				SET archived_at = NOW()
				WHERE id = $1
				  AND archived_at IS NULL;
			`

	_, delErr := database.RMS.Exec(query, sessionId)
	if delErr != nil {
		return delErr // Return error if the update fails
	}
	return nil
}

func GetAllUsersByAdmin() ([]models.User, error) {
	query := `
				SELECT u.id,
					   u.name,
					   u.email,
					   a.address,
					   ur.role
				FROM users u
						 INNER JOIN user_roles ur
									ON u.id = ur.user_id
						 INNER JOIN address a
									ON u.id = a.user_id
				WHERE u.archived_at IS NULL
				  AND ur.role = 'user';
			`

	users := make([]models.User, 0)
	FetchErr := database.RMS.Select(&users, query)
	return users, FetchErr
}

func GetAllUsersBySubAdmin(loggedUserId string) ([]models.User, error) {
	query := `
				SELECT u.id,
					   u.name,
					   u.email,
					   a.address,
					   ur.role
				FROM users u
						 INNER JOIN user_roles ur
									ON u.id = ur.user_id
						 INNER JOIN address a
									ON u.id = a.user_id
				WHERE u.archived_at IS NULL
				  AND u.created_by = $1
				  AND ur.role = 'user';
			`

	users := make([]models.User, 0)
	FetchErr := database.RMS.Select(&users, query, loggedUserId)
	return users, FetchErr
}
