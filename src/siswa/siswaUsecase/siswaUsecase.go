package siswaUsecase

import (
	"Pembayaran-Spp/model/dto/siswaDto"
	"Pembayaran-Spp/src/siswa"
)

type siswaUC struct {
	siswaRepo siswa.SiswaRepository
}

func NewSiswaUsecase(siswaRepo siswa.SiswaRepository) siswa.SiswaUsecase {
	return &siswaUC{siswaRepo}
}

func (s *siswaUC) GetLoginSiswa(nisn int) (siswaDto.SiswaResponse, error) {
	siswa, err := s.siswaRepo.LoginSiswa(nisn)
	if err != nil {
		return siswaDto.SiswaResponse{}, err
	}

	return siswa, err
}
