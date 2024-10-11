package handlers

//import (
//	"RMS/database/dbHelper"
//	"RMS/middlewares"
//	"RMS/models"
//	"RMS/utils"
//	"net/http"
//)
//
//func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
//	var body models.Restaurant
//
//	userCtx := middlewares.UserContext(r)
//	body.CreatedBy = userCtx.UserId
//
//	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
//		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
//		return
//	}
//
//	exists, existsErr := dbHelper.IsRestaurantExists(body.Name, body.CreatedBy)
//	if existsErr != nil {
//		utils.RespondError(w, http.StatusInternalServerError, existsErr, "failed to check restaurant existence")
//		return
//	}
//
//	if exists {
//		utils.RespondError(w, http.StatusConflict, nil, "restaurant already exists")
//		return
//	}
//
//	saveErr := dbHelper.CreateRestaurant(body.Name, body.Address, body.CreatedBy)
//	if saveErr != nil {
//		utils.RespondError(w, http.StatusInternalServerError, saveErr, "failed to save restaurant")
//		return
//	}
//
//	utils.RespondJSON(w, http.StatusCreated, struct {
//		Message string `json:"message"`
//	}{"restaurant created successfully"})
//}
//
//func GetAllRestaurantsByAdmin(w http.ResponseWriter, _ *http.Request) {
//	restaurants, getErr := dbHelper.GetAllRestaurantsByAdmin()
//
//	if getErr != nil {
//		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get restaurants")
//		return
//	}
//
//	if len(restaurants) == 0 {
//		utils.RespondError(w, http.StatusOK, getErr, "no restaurant found")
//		return
//	}
//
//	utils.RespondJSON(w, http.StatusOK, restaurants)
//}
//
//func GetAllRestaurantsBySubAdmin(w http.ResponseWriter, r *http.Request) {
//	userCtx := middlewares.UserContext(r)
//	loggedUserId := userCtx.UserId
//
//	restaurants, getErr := dbHelper.GetAllRestaurantsBySubAdmin(loggedUserId)
//	if getErr != nil {
//		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get restaurants")
//		return
//	}
//
//	if len(restaurants) == 0 {
//		utils.RespondError(w, http.StatusOK, getErr, "no restaurant found")
//		return
//	}
//
//	utils.RespondJSON(w, http.StatusOK, restaurants)
//}
