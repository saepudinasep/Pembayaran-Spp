package siswaDto

import (
	"Pembayaran-Spp/model/dto/kelasDto"
	"Pembayaran-Spp/model/dto/sppDto"
)

type (
	SiswaResponse struct {
		NISN   string                 `json:"nisn"`
		NIS    string                 `json:"nis"`
		Nama   string                 `json:"nama"`
		Kelas  kelasDto.KelasResponse `json:"kelas"`
		Alamat string                 `json:"alamat"`
		NoTelp string                 `json:"no_telp"`
		Spp    sppDto.SppResponse     `json:"spp"`
	}

	SiswaRequest struct {
		NISN    int    `json:"nisn"`
		NIS     int    `json:"nis"`
		Nama    string `json:"nama"`
		IdKelas int    `json:"id_kelas"`
		Alamat  string `json:"alamat"`
		NoTelp  string `json:"no_telp"`
		IdSpp   int    `json:"id_spp"`
	}
)
