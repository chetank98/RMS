package models

type CreateRestaurantRequest struct {
	Name      string  `json:"name" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type Restaurant struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Address   string  `json:"address" db:"address"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	CreatedBy string  `json:"createdBy" db:"created_by"`
}
