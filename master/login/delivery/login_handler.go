package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vivaldy22/enigma_bank/tools/vError"

	"github.com/vivaldy22/enigma_bank/models"

	"github.com/vivaldy22/enigma_bank/tools/varMux"

	"github.com/vivaldy22/enigma_bank/tools/respJson"
)

func (l *LoginHandler) ShowLogins(w http.ResponseWriter, r *http.Request) {
	data, err := l.LUseCase.GetAllLogin()

	if err != nil {
		vError.WriteError("Show Logins failed!", err, &w)
	} else {
		respJson.WriteJSON(data, w)
	}
}
func (l *LoginHandler) CreateLogin(w http.ResponseWriter, r *http.Request) {
	var login *models.Login
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		vError.WriteError("Decoding json failed!", err, &w)
	} else {
		err = l.LUseCase.Store(login)

		if err != nil {
			vError.WriteError("Create Login failed", err, &w)
		} else {
			data, err := l.LUseCase.GetByID(login.LoginID)

			if err != nil {
				vError.WriteError("Get Login by ID failed", err, &w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
func (l *LoginHandler) GetLoginByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := l.LUseCase.GetByID(idNum)

		if err != nil {
			vError.WriteError("Get Login By ID failed!", err, &w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}

}
func (l *LoginHandler) UpdateLogin(w http.ResponseWriter, r *http.Request) {
	var login *models.Login
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		vError.WriteError("Decoding json failed", err, &w)
	} else {
		id := varMux.GetVarsMux("id", r)
		idNum, err := strconv.Atoi(id)

		if err != nil {
			vError.WriteError("Converting id failed! not a number", err, &w)
		} else {
			err := l.LUseCase.Update(idNum, login)

			if err != nil {
				vError.WriteError("Updating data failed!", err, &w)
			} else {
				checkData, err := l.LUseCase.GetByID(idNum)

				if err != nil {
					vError.WriteError("Get Login By ID failed!", err, &w)
				} else {
					respJson.WriteJSON(checkData, w)
				}
			}
		}
	}
}
func (l *LoginHandler) RemoveLogin(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := l.LUseCase.GetByID(idNum)

		if err != nil {
			vError.WriteError("Get Login By ID failed!", err, &w)
		} else {
			err := l.LUseCase.Delete(idNum)

			if err != nil {
				vError.WriteError("Delete Login failed!", err, &w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
