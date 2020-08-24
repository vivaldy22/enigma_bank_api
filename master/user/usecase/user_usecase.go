package usecase

import (
	"errors"

	"github.com/vivaldy22/enigma_bank/models"
	"github.com/vivaldy22/enigma_bank/tools/validation"
)

type userUseCase struct {
	userRepo models.UserRepository
}

func (u *userUseCase) GetAllUsers() ([]*models.User, error) {
	res, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) GetByID(id int) (*models.User, error) {
	err := validation.ValidateInputNotEmpty(id)

	if err != nil {
		return nil, err
	}

	res, err := u.userRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) Store(user *models.User) error {
	err := validation.ValidateInputNotEmpty(user.LoginOwnerID, user.Balance)

	if err != nil {
		return err
	}

	if err = u.userRepo.Store(user); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Update(id int, user *models.User) error {
	err := validation.ValidateInputNotEmpty(id, user.LoginOwnerID, user.Balance)
	if err != nil {
		return err
	}

	if _, err = u.userRepo.GetByID(id); err != nil {
		return errors.New("user id not found")
	}

	if err = u.userRepo.Update(id, user); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Delete(id int) error {
	err := validation.ValidateInputNotEmpty(id)
	if err != nil {
		return err
	}

	if err = u.userRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func NewUserUseCase(repo models.UserRepository) models.UserUseCase {
	return &userUseCase{repo}
}
