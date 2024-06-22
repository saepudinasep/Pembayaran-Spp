package petugasUsecase

import (
	"Pembayaran-Spp/model/dto/petugasDto"
	"Pembayaran-Spp/src/petugas"
)

type petugasUC struct {
	petugasRepo petugas.PetugasRepository
}

func NewPetugasUsecase(petugasRepo petugas.PetugasRepository) petugas.PetugasUsecase {
	return &petugasUC{petugasRepo}
}

func (p *petugasUC) GetLogPetugas(username string) (petugasDto.PetugasResponse, error) {
	petugas, err := p.petugasRepo.LoginPetugas(username)
	if err != nil {
		return petugasDto.PetugasResponse{}, err
	}

	return petugas, err
}
