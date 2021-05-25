package service

import (
	"apiserver/pkg/model"
	"apiserver/pkg/repository"
	"errors"
	"github.com/google/uuid"
)

type UserService struct {
	repoUser repository.User
	repoRole repository.Role
}

func NewUserService(rUser repository.User, rRole repository.Role) *UserService {
	return &UserService{repoUser: rUser, repoRole: rRole}
}

func (u *UserService) GetAllUsers() ([]model.UserOutput, error) {
	return u.repoUser.GetAllUsers()
}

func (u *UserService) CreateUser(user model.User, roleName string) (uuid.UUID, error) {
	mailContatins, err := u.repoUser.CheckEmailContatinsInDB(user.Email)
	if err != nil {
		return [16]byte{}, err
	}

	if mailContatins {
		return [16]byte{}, errors.New("email is already used")
	}

	user.ID = uuid.New()
	role, err := u.repoRole.GetOneRole(roleName)
	if err != nil {
		return [16]byte{}, err
	}
	user.Role_id = role.ID
	return u.repoUser.CreateUser(user)
}

func (u *UserService) GetOneUser(user_id string) (model.UserOutput, error) {
	return u.repoUser.GetOneUser(user_id)
}

func (u *UserService) UpdateUser(user model.UpdateUserInput, user_id string) (model.UserOutput, error) {
	if err := user.Validate(); err != nil {
		return model.UserOutput{}, err
	}
	mailContatins, err := u.repoUser.CheckEmailContatinsInDB(*user.Email)
	if err != nil {
		return model.UserOutput{}, err
	}

	if mailContatins {
		return model.UserOutput{}, errors.New("email is already used")
	}

	return u.repoUser.UpdateUser(user, user_id)
}

func (u *UserService) DeleteUser(user_id string) (int64, error) {
	return u.repoUser.DeleteUser(user_id)
}
