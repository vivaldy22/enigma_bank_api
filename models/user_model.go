package models

type User struct {
	UserID       int `json:"user_id"`
	LoginOwnerID int `json:"login_owner_id"`
	Balance      int `json:"balance"`
	StatusDel    int `json:"status_del"`
}

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	GetByID(id int) (*User, error)
	Store(user *User) error
	Update(id int, user *User) error
	Delete(id int) error
}

type UserUseCase interface {
	GetAllUsers() ([]*User, error)
	GetByID(id int) (*User, error)
	Store(user *User) error
	Update(id int, user *User) error
	Delete(id int) error
}
