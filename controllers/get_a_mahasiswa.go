package controllers

import (
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/responses"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (mahasiswaController *mahasiswaController) GetAMahasiswa(c echo.Context) error {
	/*
		ASUMSI: karena soal menyatakan bahwa "mendapatkan detail data mahasiswa", karena detail berarti ada hobi juga
		STEPS:
		1. get mahasiswa by id from table friends
		2. get hobi_id from table mahasiswa_hobi by mahasiswa_id
		3. get nama_hobi from table hobi by hobi_id
	*/

	IDMahasiswa := c.QueryParam("id")

	db := mahasiswaController.db
	responseData := responses.Response[responses.MahasiswaDetail]{}
	response := responses.Response[string]{}
	detailMahasiswa := responses.MahasiswaDetail{}

	row, err := db.Query("SELECT id, nama, usia, gender, tanggal_registrasi from mahasiswa WHERE id = ?", IDMahasiswa)
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	for row.Next() {
		if err := row.Scan(&detailMahasiswa.ID, &detailMahasiswa.Nama, &detailMahasiswa.Usia, &detailMahasiswa.Gender, &detailMahasiswa.TanggalRegistrasi); err != nil {
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	defer row.Close()

	idSemuaHobi := []int{}
	rows, err := db.Query("SELECT id_hobi from mahasiswa_hobi WHERE id_mahasiswa = ?", IDMahasiswa)
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer rows.Close()
	hasResults := false
	for rows.Next() {
		hasResults = true
		var id int
		if err := rows.Scan(&id); err != nil {
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
		idSemuaHobi = append(idSemuaHobi, id)
	}
	if !hasResults {
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	semuaHobi := []string{}
	for _, idHobi := range idSemuaHobi {
		var namaHobi string
		row, err = db.Query("SELECT nama_hobi from hobi WHERE id = ?", idHobi)
		if err != nil {
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
		defer row.Close()
		for row.Next() {
			if err := row.Scan(&namaHobi); err != nil {
				response.Message = types.ERROR_INTERNAL_SERVER
				return c.JSON(http.StatusInternalServerError, response)
			}
			semuaHobi = append(semuaHobi, namaHobi)
		}
	}
	detailMahasiswa.Hobi = semuaHobi

	responseData.Data = detailMahasiswa
	responseData.Message = types.SUCCESS
	return c.JSON(http.StatusOK, responseData)
}
