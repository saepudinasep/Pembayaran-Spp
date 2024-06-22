package petugas

import "Pembayaran-Spp/model/dto/petugasDto"

type PetugasRepository interface {
	LoginPetugas(username string) (petugasDto.PetugasResponse, error)
}

type PetugasUsecase interface {
	GetLogPetugas(username string) (petugasDto.PetugasResponse, error)
}
