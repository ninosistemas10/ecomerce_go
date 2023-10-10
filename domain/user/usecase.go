package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/ecommerce/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

// User es el dominio
func (u User) CreateUsuario(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID", err)
	}
	m.ID = ID

	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "PASSWORD", err)
	}
	m.Password = string(password)

	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.CreateAt = time.Now().Unix() //Unix es para que se guarde en el horario de cualquier pais

	//Guarda en el storage la infomacion
	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf(" %s %w", "u.storage.Create", err)
	}

	m.Password = ""

	return nil

}

func (u User) GetByID(ID uuid.UUID) (model.User, error) {
	user, err := u.storage.GetByID(ID)
	if err != nil {
		return model.User{}, fmt.Errorf(" user: %w", err)
	}
	user.Password = ""
	return user, nil
}

func (u User) GetByEmaill(email string) (model.User, error) {
	users, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "storage.GetByEmail", err)
	}

	return users, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf(" %s %w", "storage.GetAll()", err)
	}
	return users, nil
}

func (u User) Login(email, password string) (model.User, error) {
	m, err := u.GetByEmaill(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "user.GetByEmail()", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "bcrypt.CompareHashAndPassword", err)
	}
	m.Password = ""
	return m, nil

}
