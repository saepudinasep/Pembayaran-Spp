package siswa

import "Pembayaran-Spp/model/dto/siswaDto"

type SiswaRepository interface {
	LoginSiswa(nisn int) (siswaDto.SiswaResponse, error)
}

type SiswaUsecase interface {
	GetLoginSiswa(nisn int) (siswaDto.SiswaResponse, error)
}
