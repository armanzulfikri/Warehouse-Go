package service

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewAuthService ...
func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &authServiceImpl{
		Repository: *userRepository,
	}
}

type authServiceImpl struct {
	Repository repository.UserRepository
}

func (service authServiceImpl) Login(request request.LoginRequest) (response response.LoginResponse) {
	user := service.Repository.GetByEmail(request.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Println("Email ", user.Email, " Password salah")
		response.Code = 401
		response.Status = "email atau password salah"
		response.Token = ""
	} else {
		type authCustomClaims struct {
			Email string `json:"email"`
			Role  string `json:"role"`
			jwt.StandardClaims
		}

		claims := &authCustomClaims{
			user.Email,
			user.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}

		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
		if err != nil {
			response.Code = 401
			response.Status = "Gagal create token, message" + err.Error()
			response.Token = ""
		} else {
			response.Code = 401
			response.Status = "Login Succcess"
			response.Token = token
		}
	}

	return response
}

func (service authServiceImpl) Register(request request.RegisterRequest) (response response.UserResponse) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := entity.Users{
		FullName:   request.FullName,
		Gender:     request.Gender,
		BirthDate:  request.BirthDate,
		BirthPlace: request.BirthPlace,
		Email:      request.FullName,
		Password:   string(hash),
		Role:       request.Role,
	}

	result := service.Repository.Insert(&user)

	if result.Email != "" {
		response.ID = result.ID
		response.Email = result.Email
		response.FullName = result.FullName
		response.Gender = result.Gender
		response.BirthDate = result.BirthDate
		response.BirthPlace = result.BirthPlace
		response.Email = result.FullName
		response.Role = result.Role
	}

	return
}
