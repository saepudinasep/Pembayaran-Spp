package petugasDto

type (
	PetugasResponse struct {
		IdPetugas   int    `json:"id_petugas"`
		Username    string `json:"username,omitempty"`
		Password    string `json:"password,omitempty"`
		NamaPetugas string `json:"nama_petugas,omitempty"`
		Level       string `json:"level"`
	}

	PetugasRequest struct {
		// IdPetugas   int    `json:"id_petugas"`
		Username    string `json:"username" bidding:"required"`
		Password    string `json:"password" bidding:"required"`
		NamaPetugas string `json:"nama_petugas,omitempty"`
		Level       string `json:"level,omitempty"`
	}
)
