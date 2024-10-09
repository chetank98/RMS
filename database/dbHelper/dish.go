package dbHelper

import (
	"RMS/database"
	"RMS/models"
)

func IsDishExists(name, restaurantId string) (bool, error) {
	query := `
				SELECT count(id) > 0 as is_exist
				FROM dishes
				WHERE name = TRIM($1)
				  AND restaurant_id = $2
				  AND archived_at IS NULL;
			`

	var check bool
	checkErr := database.RMS.Get(&check, query, name, restaurantId)
	if checkErr != nil {
		return false, checkErr
	}
	return check, nil
}

func CreateDish(name, restaurantId string, price int) error {
	query := `
				INSERT INTO dishes (name, restaurant_id, price)
				VALUES (TRIM($1), $2, $3);
			`

	_, createErr := database.RMS.Exec(query, name, restaurantId, price)
	if createErr != nil {
		return createErr
	}
	return nil
}

func GetAllDishesByAdmin() ([]models.Dish, error) {
	query := `
				SELECT d.id,
					   d.name,
					   d.price,
					   d.restaurant_id
				FROM dishes d
				WHERE d.archived_at IS NULL;
			`

	dishes := make([]models.Dish, 0)
	FetchErr := database.RMS.Select(&dishes, query)
	return dishes, FetchErr
}

func GetAllDishesBySubAdmin(loggedUserId string) ([]models.Dish, error) {
	query := `
				SELECT d.id,
					   d.name,
					   d.price,
					   d.restaurant_id
				FROM dishes d
						 INNER JOIN restaurants r on d.restaurant_id = r.id
				WHERE d.archived_at IS NULL
				  AND r.created_by = $1;
			`

	dishes := make([]models.Dish, 0)
	FetchErr := database.RMS.Select(&dishes, query, loggedUserId)
	return dishes, FetchErr
}

func DishesByRestaurant(restaurantId string) ([]models.Dish, error) {
	query := `
				SELECT d.id,
					   d.name,
					   d.price,
					   d.restaurant_id
				FROM dishes d
				WHERE d.restaurant_id = $1
				  AND d.archived_at IS NULL;
			`

	dishes := make([]models.Dish, 0)
	FetchErr := database.RMS.Select(&dishes, query, restaurantId)
	return dishes, FetchErr
}
