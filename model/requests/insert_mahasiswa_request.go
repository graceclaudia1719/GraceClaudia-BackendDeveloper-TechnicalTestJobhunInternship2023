package requests

type InsertMahasiswaRequest struct {
	Nama   string `json:"nama" form:"nama" query:"nama"`
	Usia   int    `json:"usia" form:"usia" query:"usia"`
	Gender int    `json:"gender" form:"gender" query:"gender"`
	NamaSemuaJurusan []string `json:"nama_semua_jurusan" form:"nama_semua_jurusan" query:"nama_semua_jurusan"`
	NamaSemuaHobi    []string `json:"nama_semua_hobi" form:"nama_semua_hobi" query:"nama_semua_hobi"`
}
