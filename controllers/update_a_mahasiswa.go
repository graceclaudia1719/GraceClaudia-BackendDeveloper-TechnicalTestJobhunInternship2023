package controllers

import (
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/requests"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/responses"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (mahasiswaController *mahasiswaController) UpdateAMahasiswa(c echo.Context) error {
	/*
		ASUMSI: update tidak include jurusan dan hobi karena jika include maka akan dituliskan "include jurusan dan hobi"
		seperti pada insert data mahasiswa include jurusan dan hobi

		STEPS:
			1. update pada table mahasiswa
	*/
	IDMahasiswa := c.QueryParam("id")

	db := mahasiswaController.db
	response := responses.Response[string]{}

	updateMahasiswaRequest := requests.UpdateMahasiswaRequest{}
	if err := c.Bind(&updateMahasiswaRequest); err != nil {
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}
	if !(updateMahasiswaRequest.Gender == 0 || updateMahasiswaRequest.Gender == 1) {
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	// update row
	res, err := db.Exec("UPDATE mahasiswa SET nama = ?, usia = ?, gender = ? WHERE id = ?", updateMahasiswaRequest.Nama, updateMahasiswaRequest.Usia, updateMahasiswaRequest.Gender, IDMahasiswa)
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	count, err := res.RowsAffected()
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	if count != 1 {
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}
