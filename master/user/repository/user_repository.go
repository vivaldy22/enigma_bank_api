package repository

import (
	"database/sql"

	"github.com/vivaldy22/enigma_bank/models"
	"github.com/vivaldy22/enigma_bank/tools/queries"
)

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	rows, err := u.db.Query(queries.GET_ALL_USER)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.User)
		if err := rows.Scan(&each.UserID, &each.LoginOwnerID, &each.Balance, &each.StatusDel); err != nil {
			return nil, err
		}
		users = append(users, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepo) GetByID(id int) (*models.User, error) {
	var user = new(models.User)
	row := u.db.QueryRow(queries.GET_BY_ID_LOGIN_OWNER, id)

	if err := row.Scan(&user.UserID, &user.LoginOwnerID, &user.Balance, &user.StatusDel); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) Store(user *models.User) error {
	tx, err := u.db.Begin()

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.CREATE_USER)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.LoginOwnerID, user.Balance)

	if err != nil {
		return tx.Rollback()
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return tx.Rollback()
	}

	user.UserID = int(lastInsertID)
	stmt.Close()
	return tx.Commit()
}

func (u *userRepo) Update(id int, user *models.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.UPDATE_USER)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.LoginOwnerID, user.Balance, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	user.UserID = id
	return tx.Commit()
}

func (u *userRepo) Delete(id int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.DELETE_USER)
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

func NewUserRepo(db *sql.DB) models.UserRepository {
	return &userRepo{db}
}
