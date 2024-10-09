package models

type CreateRestaurantRequest struct {
	Name      string `json:"name" db:"name"`
	Address   string `json:"address" db:"address"`
	CreatedBy string `json:"createdBy" db:"created_by"`
}

type Restaurant struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Address   string `json:"address" db:"address"`
	CreatedBy string `json:"createdBy" db:"created_by"`
}
