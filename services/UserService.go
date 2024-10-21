package services

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUser() string {
	return "{ \"name\": \"John Doe\" }"
}
