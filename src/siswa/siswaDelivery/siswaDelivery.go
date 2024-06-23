package siswaDelivery

import (
	"Pembayaran-Spp/model/dto/siswaDto"
	"Pembayaran-Spp/model/json"
	"Pembayaran-Spp/pkg/middleware"
	"Pembayaran-Spp/src/siswa"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiswaDelivery struct {
	siswaUC siswa.SiswaUsecase
}

func NewSiswaDelivery(v1Group *gin.RouterGroup, siswaUC siswa.SiswaUsecase) {
	handler := SiswaDelivery{
		siswaUC: siswaUC,
	}

	v1Group.POST("/login/siswa", handler.GetLoginSiswa).Use(middleware.BasicAuth)
}

func (s *SiswaDelivery) GetLoginSiswa(ctx *gin.Context) {
	var req siswaDto.SiswaLogReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "00", "00")
		return
	}

	siswa, err := s.siswaUC.GetLoginSiswa(req.NISN)
	if err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "00", "00")
		return
	}

	// Ensure NISN is a string before passing it to GenerateTokenJwt
	nisnStr := strconv.Itoa(siswa.NISN)

	token, err := middleware.GenerateTokenJwt(nisnStr, 3, "SISWA")
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "00", "00")
	}

	json.NewResponseSuccess(ctx, token, "success", "00", "00")
}
