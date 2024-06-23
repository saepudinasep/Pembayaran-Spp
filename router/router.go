package router

import (
	"Pembayaran-Spp/src/petugas/petugasDelivery"
	"Pembayaran-Spp/src/petugas/petugasRepository"
	"Pembayaran-Spp/src/petugas/petugasUsecase"
	"Pembayaran-Spp/src/siswa/siswaDelivery"
	"Pembayaran-Spp/src/siswa/siswaRepository"
	"Pembayaran-Spp/src/siswa/siswaUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	// Repository
	petugasRepo := petugasRepository.NewPetugasRepository(db)
	siswaRepo := siswaRepository.NewSiswaRepository(db)

	// Usecase
	petugasUC := petugasUsecase.NewPetugasUsecase(petugasRepo)
	siswaUC := siswaUsecase.NewSiswaUsecase(siswaRepo)

	// Delivery
	petugasDelivery.NewPetugasDelivery(v1Group, petugasUC)
	siswaDelivery.NewSiswaDelivery(v1Group, siswaUC)
}
