package kelasDto

type (
	KelasResponse struct {
		IdKelas            int    `json:"id_kelas"`
		NamaKelas          string `json:"nama_kelas"`
		KompetensiKeahlian string `json:"kompetensi_keahlian"`
	}

	KelasRequest struct {
		// IdKelas            int    `json:"id_kelas"`
		NamaKelas          string `json:"nama_kelas"`
		KompetensiKeahlian string `json:"kompetensi_keahlian"`
	}
)
