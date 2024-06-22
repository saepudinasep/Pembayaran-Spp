package router

import (
	"Pembayaran-Spp/src/petugas/petugasDelivery"
	"Pembayaran-Spp/src/petugas/petugasRepository"
	"Pembayaran-Spp/src/petugas/petugasUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	// Repository
	petugasRepo := petugasRepository.NewPetugasRepository(db)

	// Usecase
	petugasUC := petugasUsecase.NewPetugasUsecase(petugasRepo)

	// Delivery
	petugasDelivery.NewPetugasDelivery(v1Group, petugasUC)
}
