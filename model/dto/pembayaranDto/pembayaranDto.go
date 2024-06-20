package pembayaranDto

import (
	"Pembayaran-Spp/model/dto/petugasDto"
	"Pembayaran-Spp/model/dto/siswaDto"
	"Pembayaran-Spp/model/dto/sppDto"
)

type (
	PembayaranResponse struct {
		IdPembayaran int                        `json:"id_pembayaran"`
		Petugas      petugasDto.PetugasResponse `json:"petugas"`
		Siswa        siswaDto.SiswaResponse     `json:"siswa"`
		TglBayar     string                     `json:"tgl_bayar"`
		BulanDibayar string                     `json:"bulan_dibayar"`
		TahunDibayar string                     `json:"tahun_dibayar"`
		Spp          sppDto.SppResponse         `json:"spp"`
		JumlahBayar  int                        `json:"jumlah_bayar"`
	}

	PembayaranRequest struct {
		IdPembayaran int    `json:"id_pembayaran"`
		IdPetugas    int    `json:"id_petugas"`
		NISN         int    `json:"nisn"`
		TglBayar     string `json:"tgl_bayar"`
		BulanDibayar string `json:"bulan_dibayar"`
		TahunDibayar string `json:"tahun_dibayar"`
		IdSpp        int    `json:"id_spp"`
		JumlahBayar  int    `json:"jumlah_bayar"`
	}
)
