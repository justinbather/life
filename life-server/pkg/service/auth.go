package service

type AuthService interface {
	Authenticate(token string) (string, error)
	CreateToken(ulid string) (string, error)
}

func NewAuthService() AuthService {
	return &authService{}
}

type authService struct {
}

func (s *authService) Authenticate(token string) (string, error) {
	return "", nil
}
func (s *authService) CreateToken(ulid string) (string, error) {
	return "", nil
}
