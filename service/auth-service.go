package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//AuthService ...
type AuthService interface {
	Login(request request.LoginRequest) (response response.LoginResponse)
	Register(request request.RegisterRequest) (response response.UserResponse)
}
