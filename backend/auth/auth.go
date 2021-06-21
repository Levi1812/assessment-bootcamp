package auth

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type Service interface {
	GenerateToken(UserId int) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(UserId int) (string, error) {

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

}
