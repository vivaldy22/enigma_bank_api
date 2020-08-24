package usecase

import (
	"errors"

	"github.com/vivaldy22/enigma_bank/models"
	"github.com/vivaldy22/enigma_bank/tools/validation"
)

type transactionUseCase struct {
	transactionRepo models.TransactionRepository
}

func (u *transactionUseCase) GetAllTransactions() ([]*models.Transaction, error) {
	res, err := u.transactionRepo.GetAllTransactions()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUseCase) GetByID(id int) (*models.Transaction, error) {
	err := validation.ValidateInputNotEmpty(id)

	if err != nil {
		return nil, err
	}

	res, err := u.transactionRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUseCase) GetByUserOwnerID(id int) ([]*models.Transaction, error) {
	err := validation.ValidateInputNotEmpty(id)

	if err != nil {
		return nil, err
	}

	res, err := u.transactionRepo.GetByUserOwnerID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUseCase) Store(transaction *models.Transaction) error {
	err := validation.ValidateInputNotEmpty(transaction.UserOwnerID, transaction.TransDate, transaction.Destination,
		transaction.Amount, transaction.Description)

	if err != nil {
		return err
	}

	if err = u.transactionRepo.Store(transaction); err != nil {
		return err
	}

	return nil
}

func (u *transactionUseCase) Update(id int, transaction *models.Transaction) error {
	err := validation.ValidateInputNotEmpty(id, transaction.UserOwnerID, transaction.TransDate, transaction.Destination,
		transaction.Amount, transaction.Description)
	if err != nil {
		return err
	}

	if _, err = u.transactionRepo.GetByID(id); err != nil {
		return errors.New("transaction id not found")
	}

	if err = u.transactionRepo.Update(id, transaction); err != nil {
		return err
	}

	return nil
}

func (u *transactionUseCase) Delete(id int) error {
	err := validation.ValidateInputNotEmpty(id)
	if err != nil {
		return err
	}

	if err = u.transactionRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func NewTransactionUseCase(repo models.TransactionRepository) models.TransactionUseCase {
	return &transactionUseCase{repo}
}
