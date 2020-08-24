package models

type Transaction struct {
	TransID     int    `json:"trans_id"`
	UserOwnerID int    `json:"user_owner_id"`
	TransDate   string `json:"trans_date"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	StatusDel   int    `json:"-"`
}

type TransactionRepository interface {
	GetAllTransactions() ([]*Transaction, error)
	GetByID(id int) (*Transaction, error)
	Store(trans *Transaction) error
	Update(id int, movie *Transaction) error
	Delete(id int) error
}

type TransactionUseCase interface {
	GetAllTransactions() ([]*Transaction, error)
	GetByID(id int) (*Transaction, error)
	Store(trans *Transaction) error
	Update(id int, movie *Transaction) error
	Delete(id int) error
}
