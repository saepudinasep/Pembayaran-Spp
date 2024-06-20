package authRepository

import (
	"Pembayaran-Spp/model/dto/petugasDto"
	"Pembayaran-Spp/src/auth"
	"database/sql"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) auth.AuthRepository {
	return &authRepository{db}
}

func (a *authRepository) RetrieveLogin(username string) (petugasDto.PetugasResponse, error) {
	// query untuk mengambil data user berdasarkan username
	var petugas petugasDto.PetugasResponse
	err := a.db.QueryRow("SELECT id_petugas, username, password, nama_petugas, level FROM tb_petugas WHERE username = $1", username).Scan(&petugas.IdPetugas, &petugas.Username, &petugas.Password, &petugas.NamaPetugas, &petugas.Level)
	if err != nil {
		return petugasDto.PetugasResponse{}, err
	}

	return petugas, err
}
