package handlers

import (
	"RMS/database"
	"RMS/database/dbHelper"
	"RMS/middlewares"
	"RMS/models"
	"RMS/utils"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CreateSubAdmin(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterSubAdminRequest

	userCtx := middlewares.UserContext(r)
	body.CreatedBy = userCtx.UserId

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	if !body.Role.IsValid() {
		utils.RespondError(w, http.StatusBadRequest, nil, "invalid role type provided")
		return
	}

	exists, existsErr := dbHelper.IsUserExists(body.Email)
	if existsErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, existsErr, "failed to check user existence")
		return
	}

	if exists {
		utils.RespondError(w, http.StatusConflict, nil, "sub-admin already exists")
		return
	}

	hashedPassword, hasErr := utils.HashPassword(body.Password)
	if hasErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, hasErr, "failed to secure password")
		return
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		userID, saveErr := dbHelper.CreateUser(tx, body.Name, body.Email, hashedPassword, body.CreatedBy)
		if saveErr != nil {
			utils.RespondError(w, http.StatusInternalServerError, saveErr, "failed to save sub-admin")
			return saveErr
		}
		roleErr := dbHelper.CreateUserRole(tx, userID, body.Role)
		if roleErr != nil {
			return roleErr
		}
		return nil
	})
	if txErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, txErr, "failed to create sub-admin")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
	}{"sub-admin created successfully"})
}

func GetAllSubAdmins(w http.ResponseWriter, _ *http.Request) {
	subAdmins, getErr := dbHelper.GetAllSubAdmins()

	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get sub-admin")
		return
	}

	if len(subAdmins) == 0 {
		utils.RespondError(w, http.StatusOK, getErr, "no sub-admin found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, subAdmins)
}
