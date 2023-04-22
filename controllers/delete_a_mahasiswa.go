package controllers

import (
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/responses"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (mahasiswaController *mahasiswaController) DeleteAMahasiswa(c echo.Context) error {
	/*
		STEPS (using transaction):
			1. delete if exist in mahasiswa table
			2. delete all occurences of mahasiswa_id in mahasiswa_hobi table

		note: tidak perlu mendelet hobi pada hobi table karena bisa saja ada kasus dimana mahasiswa lain
		memiliki hobi yang sama. Jika harus di delete, maka perlu mengiterasi ke semua row pada mahasiswa_hobi
		dan mengecek apakah hobi_id yang akan didelet ada pada tabel tersebut, jika tidak maka akan didelet di tabel
		hobi. Namun, operasi ini cukup mahal karena perlu mengiterasi semua row pada tabel mahasiswa_hobi sehingga akan lebih
		baik jika tidak didelet.

	*/

	IDMahasiswa := c.QueryParam("id")
	db := mahasiswaController.db

	response := responses.Response[string]{}

	// TRANSACTION
	tx, err := db.Begin()
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// delete in mahasiswa_hobi
	res, err := tx.Exec("DELETE FROM mahasiswa_hobi WHERE id_mahasiswa = ?", IDMahasiswa)
	if err != nil {
		tx.Rollback()
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	count, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	if count < 0 {
		tx.Rollback()
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	// delete in mahasiswa
	res, err = tx.Exec("DELETE FROM mahasiswa WHERE id = ?", IDMahasiswa)
	if err != nil {
		tx.Rollback()
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	count, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	if count != 1 {
		tx.Rollback()
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}
	
	tx.Commit()

	response.Message = types.SUCCESS
	return c.JSON(http.StatusAccepted, response)
}
