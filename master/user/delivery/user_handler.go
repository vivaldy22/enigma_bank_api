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

func (u *UserHandler) ShowUsers(w http.ResponseWriter, r *http.Request) {
	data, err := u.LUseCase.GetAllUsers()

	if err != nil {
		vError.WriteError("Show Users failed!", err, &w)
	} else {
		respJson.WriteJSON(data, w)
	}
}
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", err, &w)
	} else {
		err = u.LUseCase.Store(user)

		if err != nil {
			vError.WriteError("Create User failed", err, &w)
		} else {
			data, err := u.LUseCase.GetByID(user.UserID)

			if err != nil {
				vError.WriteError("Get User by ID failed", err, &w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := u.LUseCase.GetByID(idNum)

		if err != nil {
			vError.WriteError("Get User By ID failed!", err, &w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}

}
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed", err, &w)
	} else {
		id := varMux.GetVarsMux("id", r)
		idNum, err := strconv.Atoi(id)

		if err != nil {
			vError.WriteError("Converting id failed! not a number", err, &w)
		} else {
			err := u.LUseCase.Update(idNum, user)

			if err != nil {
				vError.WriteError("Updating data failed!", err, &w)
			} else {
				checkData, err := u.LUseCase.GetByID(idNum)

				if err != nil {
					vError.WriteError("Get User By ID failed!", err, &w)
				} else {
					respJson.WriteJSON(checkData, w)
				}
			}
		}
	}
}
func (u *UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := u.LUseCase.GetByID(idNum)

		if err != nil {
			vError.WriteError("Get User By ID failed!", err, &w)
		} else {
			err := u.LUseCase.Delete(idNum)

			if err != nil {
				vError.WriteError("Delete User failed!", err, &w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
