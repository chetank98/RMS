package models

type CreateDishRequest struct {
	Name         string `json:"name" db:"name"`
	Price        int    `json:"price" db:"price"`
	RestaurantId string `json:"restaurantId" db:"restaurant_id"`
}

type Dish struct {
	Id           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Price        int    `json:"price" db:"price"`
	RestaurantId string `json:"restaurantId" db:"restaurant_id"`
}
