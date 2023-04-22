package controllers

import (
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/responses"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (mahasiswaController *mahasiswaController) GetAllMahasiswa(c echo.Context) error {
	/*
		ASUMSI: semua data mahasiswa berarti data yang ada pada entitas mahasiswa saja, untuk detail
		baru dikeluarkan hobinya.

		STEPS:
			1. get all mahasiswa from mahasiswa table
	*/

	db := mahasiswaController.db
	response := responses.Response[[]responses.Mahasiswa]{}

	rows, err := db.Query("SELECT id, nama, usia, gender, tanggal_registrasi from mahasiswa")
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer rows.Close()

	for rows.Next() {
		dataMahasiswa := responses.Mahasiswa{}
		if err := rows.Scan(
			&dataMahasiswa.ID,
			&dataMahasiswa.Nama,
			&dataMahasiswa.Usia,
			&dataMahasiswa.Gender,
			&dataMahasiswa.TanggalRegistrasi); err != nil {
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
		response.Data = append(response.Data, dataMahasiswa)
	}
	if err = rows.Err(); err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}
