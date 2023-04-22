package requests

type UpdateMahasiswaRequest struct {
	Nama   string `json:"nama" form:"nama" query:"nama"`
	Usia   int    `json:"usia" form:"usia" query:"usia"`
	Gender int    `json:"gender" form:"gender" query:"gender"`
}
