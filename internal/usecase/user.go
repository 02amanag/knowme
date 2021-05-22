package usecase

import (
	"errors"

	"github.com/02amanag/p-02/internal/helper"
	"github.com/02amanag/p-02/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (u *UsecaseStruct) Login(email, password string) (*model.TokenDetails, string, error) {
	user := u.repo.GetUsers(email)
	if user.Email != email {
		return nil, "", errors.New("Invalid Email")
	}
	ok := helper.MatchByteEncryptPassword(password, user.Password)
	if ok {
		// if user.Password == password {
		username := u.repo.GetUsername(int(user.Id))
		token, err := u.CreateToken(user.Id, username)
		if err != nil {
			return nil, "", err
		}
		return token, username, nil
	}

	return nil, "", errors.New("Invalid password")

}

func (u *UsecaseStruct) SingUp(email, password, fname, lname string) error {
	user := u.repo.GetUsers(email)

	if user.Email == email {
		return errors.New("Email alredy registered")
	}
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = u.repo.CreateUser(email, string(hashedPassword), fname, lname)
	if err != nil {
		return err
	}
	return nil
}
