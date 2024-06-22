package petugasRepository

import (
	"Pembayaran-Spp/model/dto/petugasDto"
	"Pembayaran-Spp/src/petugas"
	"database/sql"
)

type petugasRepository struct {
	db *sql.DB
}

func NewPetugasRepository(db *sql.DB) petugas.PetugasRepository {
	return &petugasRepository{db}
}

func (p *petugasRepository) LoginPetugas(username string) (petugasDto.PetugasResponse, error) {
	var petugas petugasDto.PetugasResponse
	err := p.db.QueryRow("SELECT id_petugas, username, password, nama_petugas, level FROM tb_petugas WHERE username = $1", username).Scan(&petugas.IdPetugas, &petugas.Username, &petugas.Password, &petugas.NamaPetugas, &petugas.Level)
	if err != nil {
		return petugasDto.PetugasResponse{}, err
	}

	return petugas, err
}
