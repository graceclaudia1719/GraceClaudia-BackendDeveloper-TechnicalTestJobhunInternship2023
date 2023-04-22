package model

import "time"

type Mahasiswa struct {
	ID                int       `json:"id" form:"id" query:"id"`
	Nama              string    `json:"nama" form:"nama" query:"nama"`
	Usia              int       `json:"usia" form:"usia" query:"usia"`
	Gender            int       `json:"gender" form:"gender" query:"gender"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi" form:"tanggal_registrasi" query:"tanggal_registrasi"`
}
