package dbHelper

import (
	"RMS/database"
	"RMS/models"
)

func IsRestaurantExists(name, address string) (bool, error) {
	SQL := ` SELECT count(id) > 0 as is_exist
				FROM restaurants
				WHERE name = TRIM($1)
				  AND address = TRIM($2)
				  AND archived_at IS NULL`

	var check bool
	checkErr := database.RMS.Get(&check, SQL, name, address)
	return check, checkErr
}

func CreateRestaurant(body models.CreateRestaurantRequest, userID string) error {
	args := []interface{}{body.Name, body.Address, body.Latitude, body.Longitude, userID}

	SQL := `INSERT INTO restaurants (name, address, latitude, longitude, created_by)
				VALUES (TRIM($1), TRIM($2), $3, $4, $5)`

	_, createErr := database.RMS.Exec(SQL, args...)
	return createErr
}

func GetAllRestaurants() ([]models.Restaurant, error) {
	SQL := `SELECT id, name, address,
				   latitude, longitude, created_by
			  FROM restaurants
				WHERE archived_at IS NULL`

	restaurants := make([]models.Restaurant, 0)
	fetchErr := database.RMS.Select(&restaurants, SQL)
	return restaurants, fetchErr
}

func GetAllRestaurantsBySubAdmin(loggedUserID string) ([]models.Restaurant, error) {
	SQL := `SELECT id, name, address,
				   latitude, longitude, created_by
			  FROM restaurants
				WHERE created_by = $1
				    AND archived_at IS NULL`

	restaurants := make([]models.Restaurant, 0)
	fetchErr := database.RMS.Select(&restaurants, SQL, loggedUserID)
	return restaurants, fetchErr
}
