package handlers

import (
	"RMS/database/dbHelper"
	"RMS/middlewares"
	"RMS/models"
	"RMS/utils"
	"net/http"
)

func CreateDish(w http.ResponseWriter, r *http.Request) {
	var body models.Dish

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	exists, existsErr := dbHelper.IsDishExists(body.Name, body.RestaurantId)
	if existsErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, existsErr, "failed to check dish existence")
		return
	}
	if exists {
		utils.RespondError(w, http.StatusConflict, nil, "dish already exists")
		return
	}

	saveErr := dbHelper.CreateDish(body.Name, body.RestaurantId, body.Price)
	if saveErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, saveErr, "failed to save dish")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
	}{"dish created successfully"})
}

func GetAllDishesByAdmin(w http.ResponseWriter, _ *http.Request) {
	dishes, getErr := dbHelper.GetAllDishesByAdmin()

	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get dishes")
		return
	}

	if len(dishes) == 0 {
		utils.RespondError(w, http.StatusOK, getErr, "no dish found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, dishes)
}

func GetAllDishesBySubAdmin(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.UserContext(r)
	loggedUserId := userCtx.UserId

	dishes, getErr := dbHelper.GetAllDishesBySubAdmin(loggedUserId)
	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get dishes")
		return
	}

	if len(dishes) == 0 {
		utils.RespondError(w, http.StatusOK, getErr, "no dish found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, dishes)
}

func DishesByRestaurant(w http.ResponseWriter, r *http.Request) {
	body := struct {
		RestaurantId string `json:"restaurantId" db:"restaurant_id"`
	}{}

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	dishes, getErr := dbHelper.DishesByRestaurant(body.RestaurantId)
	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get dishes")
		return
	}

	if len(dishes) == 0 {
		utils.RespondError(w, http.StatusOK, getErr, "no dish found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, dishes)
}
