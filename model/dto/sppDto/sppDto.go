package sppDto

type (
	SppResponse struct {
		IdSpp   int `json:"id_spp"`
		Tahun   int `json:"tahun"`
		Nominal int `json:"nominal"`
	}

	SppRequest struct {
		// IdSpp   int `json:"id_spp"`
		Tahun   int `json:"tahun"`
		Nominal int `json:"nominal"`
	}
)
