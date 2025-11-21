package domain

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Save(user *User) (*User, error)
	GetByID(id int) (*User, error)
	GetAll() ([]User, error)
}

type UserUseCase interface {
	CreateUser(user *User) (*User, error)
	GetByID(id int) (*User, error)
	GetAll() ([]User, error)
}
