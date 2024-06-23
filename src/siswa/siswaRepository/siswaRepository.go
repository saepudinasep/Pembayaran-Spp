package siswaRepository

import (
	"Pembayaran-Spp/model/dto/siswaDto"
	"Pembayaran-Spp/src/siswa"
	"database/sql"
)

type siswaRepository struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) siswa.SiswaRepository {
	return &siswaRepository{db}
}

func (s *siswaRepository) LoginSiswa(nisn int) (siswaDto.SiswaResponse, error) {
	var siswa siswaDto.SiswaResponse
	err := s.db.QueryRow("SELECT s.nisn, s.nis, s.nama, k.id_kelas, k.nama_kelas, k.kompetensi_keahlian, s.alamat, s.no_telp, sp.id_spp, sp.tahun, sp.nominal FROM tb_siswa s JOIN tb_kelas k ON s.id_kelas = k.id_kelas JOIN tb_spp sp ON s.id_spp = sp.id_spp WHERE s.nisn = $1", nisn).Scan(
		&siswa.NISN,
		&siswa.NIS,
		&siswa.Nama,
		&siswa.Kelas.IdKelas,
		&siswa.Kelas.NamaKelas,
		&siswa.Kelas.KompetensiKeahlian,
		&siswa.Alamat,
		&siswa.NoTelp,
		&siswa.Spp.IdSpp,
		&siswa.Spp.Tahun,
		&siswa.Spp.Nominal,
	)
	if err != nil {
		return siswaDto.SiswaResponse{}, err
	}

	return siswa, err
}
