package petugasDelivery

import (
	"Pembayaran-Spp/model/dto/petugasDto"
	"Pembayaran-Spp/model/json"
	"Pembayaran-Spp/pkg/middleware"
	"Pembayaran-Spp/src/petugas"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PetugasDelivery struct {
	petugasUC petugas.PetugasUsecase
}

func NewPetugasDelivery(v1Group *gin.RouterGroup, petugasUC petugas.PetugasUsecase) {
	handler := PetugasDelivery{
		petugasUC: petugasUC,
	}

	v1Group.POST("/login/petugas", handler.GetLogPetugas).Use(middleware.BasicAuth)

	petugasGroup := v1Group.Group("/petugas")
	petugasGroup.Use(middleware.JWTAuth("ADMIN"))
	{
		petugasGroup.POST("/", handler.GetLogPetugas)
	}
}

func (p *PetugasDelivery) GetLogPetugas(ctx *gin.Context) {
	var req petugasDto.PetugasLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "00", "00")
		return
	}

	petugas, err := p.petugasUC.GetLogPetugas(req.Username)
	if err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "00", "00")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(petugas.Password), []byte(req.Password))
	if err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "Invalid hash password", "01", "01")
		return
	}

	token, err := middleware.GenerateTokenJwt(req.Username, 3, petugas.Level)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "00", "00")
		return
	}

	json.NewResponseSuccess(ctx, token, "success", "01", "01")
}
