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

func (u *TransactionHandler) ShowTransactions(w http.ResponseWriter, r *http.Request) {
	data, err := u.TUseCase.GetAllTransactions()

	if err != nil {
		vError.WriteError("Show Transactions failed!", err, &w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (u *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction *models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		vError.WriteError("Decoding json failed!", err, &w)
	} else {
		err = u.TUseCase.Store(transaction)

		if err != nil {
			vError.WriteError("Create Transaction failed", err, &w)
		} else {
			data, err := u.TUseCase.GetByID(transaction.TransID)

			if err != nil {
				vError.WriteError("Get Transaction by ID failed", err, &w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (u *TransactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := u.TUseCase.GetByID(idNum)

		if err != nil {
			vError.WriteError("Get Transaction By ID failed!", err, &w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}

func (u *TransactionHandler) GetTransactionByUserOwnerID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := u.TUseCase.GetByUserOwnerID(idNum)

		if err != nil {
			vError.WriteError("Get Transaction By User Owner ID failed!", err, &w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}

func (u *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction *models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		vError.WriteError("Decoding json failed", err, &w)
	} else {
		id := varMux.GetVarsMux("id", r)
		idNum, err := strconv.Atoi(id)

		if err != nil {
			vError.WriteError("Converting id failed! not a number", err, &w)
		} else {
			err := u.TUseCase.Update(idNum, transaction)

			if err != nil {
				vError.WriteError("Updating data failed!", err, &w)
			} else {
				checkData, err := u.TUseCase.GetByID(idNum)

				if err != nil {
					vError.WriteError("Get Transaction By ID failed!", err, &w)
				} else {
					respJson.WriteJSON(checkData, w)
				}
			}
		}
	}
}

func (u *TransactionHandler) RemoveTransaction(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", err, &w)
	} else {
		data, err := u.TUseCase.GetByID(idNum)

		if err != nil {
			vError.WriteError("Get Transaction By ID failed!", err, &w)
		} else {
			err := u.TUseCase.Delete(idNum)

			if err != nil {
				vError.WriteError("Delete Transaction failed!", err, &w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
