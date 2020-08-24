package models

type Login struct {
	LoginID   int    `json:"login_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	StatusDel int    `json:"-"`
}

type LoginRepository interface {
	GetAllLogin() ([]*Login, error)
	GetByID(id int) (*Login, error)
	Store(login *Login) error
	Update(id int, login *Login) error
	Delete(id int) error
}

type LoginUseCase interface {
	GetAllLogin() ([]*Login, error)
	GetByID(id int) (*Login, error)
	Store(login *Login) error
	Update(id int, login *Login) error
	Delete(id int) error
}
