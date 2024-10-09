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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterUserRequest

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
		utils.RespondError(w, http.StatusConflict, nil, "user already exists")
		return
	}

	hashedPassword, hasErr := utils.HashPassword(body.Password)
	if hasErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, hasErr, "failed to secure password")
		return
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		userId, saveErr := dbHelper.CreateUser(tx, body.Name, body.Email, hashedPassword, body.CreatedBy)
		if saveErr != nil {
			return saveErr
		}
		addErr := dbHelper.CreateUserAddress(tx, userId, body.Address)
		if addErr != nil {
			return addErr
		}
		roleErr := dbHelper.CreateUserRole(tx, userId, body.Role)
		if roleErr != nil {
			return roleErr
		}
		return nil
	})
	if txErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, txErr, "failed to create user")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
	}{"user created successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginUser

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	var userInfo *models.UserInfo
	userInfo, getErr := dbHelper.GetUserInfo(body.Email, body.Password)
	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to find user")
		return
	}

	if userInfo == nil {
		utils.RespondError(w, http.StatusOK, nil, "user not found")
		return
	}

	sessionId, crtErr := dbHelper.CreateUserSession(userInfo.Id)
	if crtErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, crtErr, "failed to create user session")
		return
	}

	token, genErr := utils.GenerateJWT(userInfo.Id, sessionId, userInfo.Role)
	if genErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, genErr, "failed to generate token")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{"user logged in successfully", token})
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.UserContext(r)
	sessionId := userCtx.SessionId

	saveErr := dbHelper.DeleteUserSession(sessionId)
	if saveErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, saveErr, "failed to delete user session")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{"user logged out successfully"})
}

func GetAllUsersByAdmin(w http.ResponseWriter, _ *http.Request) {
	users, getErr := dbHelper.GetAllUsersByAdmin()

	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get users")
		return
	}

	if len(users) == 0 {
		utils.RespondError(w, http.StatusOK, getErr, "no user found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, users)
}

func GetAllUsersBySubAdmin(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.UserContext(r)
	loggedUserId := userCtx.UserId

	users, getErr := dbHelper.GetAllUsersBySubAdmin(loggedUserId)
	if getErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, getErr, "failed to get users")
		return
	}

	if len(users) == 0 {
		utils.RespondError(w, http.StatusOK, getErr, "no user found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, users)
}
