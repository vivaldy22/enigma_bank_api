package repository

import (
	"database/sql"

	"github.com/vivaldy22/enigma_bank/models"
	"github.com/vivaldy22/enigma_bank/tools/queries"
)

type transactionRepo struct {
	db *sql.DB
}

func (u *transactionRepo) GetAllTransactions() ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	rows, err := u.db.Query(queries.GET_ALL_TRANSACTION)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Transaction)
		if err := rows.Scan(&each.TransID, &each.UserOwnerID, &each.TransDate, &each.Destination, &each.Amount,
			&each.Description, &each.StatusDel); err != nil {
			return nil, err
		}
		transactions = append(transactions, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (u *transactionRepo) GetByID(id int) (*models.Transaction, error) {
	var transaction = new(models.Transaction)
	row := u.db.QueryRow(queries.GET_BY_ID_TRANSACTION, id)

	if err := row.Scan(&transaction.TransID, &transaction.UserOwnerID, &transaction.TransDate, &transaction.Destination,
		&transaction.Amount, &transaction.Description, &transaction.StatusDel); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (u *transactionRepo) GetByUserOwnerID(id int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	rows, err := u.db.Query(queries.GET_BY_ID_USER_OWNER, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Transaction)
		if err := rows.Scan(&each.TransID, &each.UserOwnerID, &each.TransDate, &each.Destination, &each.Amount,
			&each.Description, &each.StatusDel); err != nil {
			return nil, err
		}
		transactions = append(transactions, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (u *transactionRepo) Store(transaction *models.Transaction) error {
	tx, err := u.db.Begin()

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.CREATE_TRANSACTION)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(transaction.UserOwnerID, transaction.TransDate, transaction.Destination,
		transaction.Amount, transaction.Description)

	if err != nil {
		return tx.Rollback()
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return tx.Rollback()
	}

	transaction.TransID = int(lastInsertID)
	stmt.Close()
	return tx.Commit()
}

func (u *transactionRepo) Update(id int, transaction *models.Transaction) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.UPDATE_TRANSACTION)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(transaction.TransDate, transaction.Destination,
		transaction.Amount, transaction.Description, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	transaction.TransID = id
	return tx.Commit()
}

func (u *transactionRepo) Delete(id int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.DELETE_TRANSACTION)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func NewTransactionRepo(db *sql.DB) models.TransactionRepository {
	return &transactionRepo{db}
}
