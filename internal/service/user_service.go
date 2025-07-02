package service

import (
	"github.com/Shashanktriathi1703/student-api/internal/model"
	"github.com/Shashanktriathi1703/student-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo : repo}
}

func (s *UserService) CreatedUser(user *model.CreatedUserRequest)(*model.User, error){
	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id int)(*model.User, error){
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers()([]model.User, error){
	return s.repo.GetAll()
}

func (s *UserService) DeleteUserid(id int)error{
	return s.repo.Delete(id)
}