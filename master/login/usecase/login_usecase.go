package usecase

import (
	"errors"

	"github.com/vivaldy22/enigma_bank/models"
	"github.com/vivaldy22/enigma_bank/tools/validation"
)

type loginUseCase struct {
	loginRepo models.LoginRepository
}

func (l *loginUseCase) GetAllLogin() ([]*models.Login, error) {
	res, err := l.loginRepo.GetAllLogin()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *loginUseCase) GetByID(id int) (*models.Login, error) {
	err := validation.ValidateInputNotEmpty(id)

	if err != nil {
		return nil, err
	}

	res, err := l.loginRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *loginUseCase) Store(login *models.Login) error {
	err := validation.ValidateInputNotEmpty(login.Username, login.Password)

	if err != nil {
		return err
	}

	if err = l.loginRepo.Store(login); err != nil {
		return err
	}

	return nil
}

func (l *loginUseCase) Update(id int, login *models.Login) error {
	err := validation.ValidateInputNotEmpty(id, login.Username, login.Password)
	if err != nil {
		return err
	}

	if _, err = l.loginRepo.GetByID(id); err != nil {
		return errors.New("login id not found")
	}

	if err = l.loginRepo.Update(id, login); err != nil {
		return err
	}

	return nil
}

func (l *loginUseCase) Delete(id int) error {
	err := validation.ValidateInputNotEmpty(id)
	if err != nil {
		return err
	}

	if err = l.loginRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func NewLoginUseCase(repo models.LoginRepository) models.LoginUseCase {
	return &loginUseCase{repo}
}
