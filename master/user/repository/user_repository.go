package repository

import (
	"database/sql"

	"github.com/vivaldy22/enigma_bank/models"
	"github.com/vivaldy22/enigma_bank/tools/queries"
)

type loginRepo struct {
	db *sql.DB
}

func (l *loginRepo) GetAllLogin() ([]*models.Login, error) {
	var logins []*models.Login
	rows, err := l.db.Query(queries.GET_ALL_LOGIN)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Login)
		if err := rows.Scan(&each.LoginID, &each.Username, &each.Password, &each.StatusDel); err != nil {
			return nil, err
		}
		logins = append(logins, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return logins, nil
}

func (l *loginRepo) GetByID(id int) (*models.Login, error) {
	var login = new(models.Login)
	row := l.db.QueryRow(queries.GET_BY_ID_LOGIN, id)

	if err := row.Scan(&login.LoginID, &login.Username, &login.Password, &login.StatusDel); err != nil {
		return nil, err
	}
	return login, nil
}

func (l *loginRepo) Store(login *models.Login) error {
	tx, err := l.db.Begin()

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.CREATE_LOGIN)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(login.Username, login.Password)

	if err != nil {
		return tx.Rollback()
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return tx.Rollback()
	}

	login.LoginID = int(lastInsertID)
	stmt.Close()
	return tx.Commit()
}

func (l *loginRepo) Update(id int, login *models.Login) error {
	tx, err := l.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.UPDATE_LOGIN)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(login.Username, login.Password, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	login.LoginID = id
	return tx.Commit()
}

func (l *loginRepo) Delete(id int) error {
	tx, err := l.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.DELETE_LOGIN)
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

func NewLoginRepo(db *sql.DB) models.LoginRepository {
	return &loginRepo{db}
}
