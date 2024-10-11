package dbHelper

//
//import (
//	"RMS/database"
//	"RMS/models"
//)
//
//func IsRestaurantExists(name, userId string) (bool, error) {
//	query := `
//				SELECT count(id) > 0 as is_exist
//				FROM restaurants
//				WHERE name = TRIM($1)
//				  AND created_by = $2
//				  AND archived_at IS NULL;
//			`
//
//	var check bool
//	checkErr := database.RMS.Get(&check, query, name, userId)
//	if checkErr != nil {
//		return false, checkErr
//	}
//	return check, nil
//}
//
//func CreateRestaurant(name, address, userId string) error {
//	query := `
//				INSERT INTO restaurants (name, address, created_by)
//				VALUES (TRIM($1), TRIM($2), $3);
//			`
//
//	_, createErr := database.RMS.Exec(query, name, address, userId)
//	if createErr != nil {
//		return createErr
//	}
//	return nil
//}
//
//func GetAllRestaurantsByAdmin() ([]models.Restaurant, error) {
//	query := `
//				SELECT r.id,
//					   r.name,
//					   r.address,
//					   r.created_by
//				FROM restaurants r
//				WHERE r.archived_at IS NULL;
//			`
//
//	restaurants := make([]models.Restaurant, 0)
//	fetchErr := database.RMS.Select(&restaurants, query)
//	return restaurants, fetchErr
//}
//
//func GetAllRestaurantsBySubAdmin(loggedUserId string) ([]models.Restaurant, error) {
//	query := `
//				SELECT r.id,
//					   r.name,
//					   r.address,
//					   r.created_by
//				FROM restaurants r
//				WHERE r.archived_at IS NULL
//				  AND created_by = $1;
//			`
//
//	restaurants := make([]models.Restaurant, 0)
//	fetchErr := database.RMS.Select(&restaurants, query, loggedUserId)
//	return restaurants, fetchErr
//}
