package auth

import (
	"Pembayaran-Spp/model/dto/petugasDto"
)

type AuthRepository interface {
	RetrieveLogin(username string) (petugasDto.PetugasResponse, error)
}

type AuthUsecase interface {
	Login(petugas petugasDto.PetugasRequest) (token string, err error)
}
