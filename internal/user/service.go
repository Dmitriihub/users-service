package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateUserByID(id uint, user *User) (*User, error) {
	return s.repo.UpdateByID(id, user)
}

func (s *Service) DeleteUserByID(id uint) error {
	return s.repo.DeleteByID(id)
}
