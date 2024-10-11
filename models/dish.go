package models

type CreateDishRequest struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
}

type Dish struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Price        int    `json:"price" db:"price"`
	RestaurantID string `json:"restaurantId" db:"restaurant_id"`
}
