package dbHelper

import (
	"RMS/database"
	"RMS/models"
)

func IsDishExists(name, restaurantID string) (bool, error) {
	SQL := `SELECT count(id) > 0 as is_exist
				FROM dishes
				WHERE name = TRIM($1)
				  AND restaurant_id = $2
				  AND archived_at IS NULL`

	var check bool
	checkErr := database.RMS.Get(&check, SQL, name, restaurantID)
	return check, checkErr
}

func CreateDish(body models.CreateDishRequest, restaurantID string) error {
	SQL := `INSERT INTO dishes (name, price, restaurant_id)
				VALUES (TRIM($1), $2, $3)`

	_, createErr := database.RMS.Exec(SQL, body.Name, body.Price, restaurantID)
	return createErr
}

func GetAllDishes() ([]models.Dish, error) {
	SQL := `SELECT id, name, price, restaurant_id
				FROM dishes
				WHERE archived_at IS NULL`

	dishes := make([]models.Dish, 0)
	FetchErr := database.RMS.Select(&dishes, SQL)
	return dishes, FetchErr
}

func GetAllDishesBySubAdmin(loggedUserID string) ([]models.Dish, error) {
	SQL := `SELECT d.id, d.name, d.price, d.restaurant_id
				FROM dishes d
						 INNER JOIN restaurants r on d.restaurant_id = r.id
				WHERE d.archived_at IS NULL
				  AND r.created_by = $1`

	dishes := make([]models.Dish, 0)
	fetchErr := database.RMS.Select(&dishes, SQL, loggedUserID)
	return dishes, fetchErr
}

func DishesByRestaurant(restaurantID string) ([]models.Dish, error) {
	SQL := `SELECT id, name, price, restaurant_id
				FROM dishes
				WHERE restaurant_id = $1
				  AND archived_at IS NULL`

	dishes := make([]models.Dish, 0)
	fetchErr := database.RMS.Select(&dishes, SQL, restaurantID)
	return dishes, fetchErr
}
