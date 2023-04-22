package controllers

import (
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/requests"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/responses"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/types"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (mahasiswaController *mahasiswaController) InsertANewMahasiswa(c echo.Context) error {

	/*
		STEP (using transaction):
		1. insert mahasiswa struct to mahasiswa table
		2. insert hobi if not exist to hobi table
		3. insert jurusan if not exist to jurusan table
		4. insert mahasiswa_id and hobi_id to mahasiswa_hobi table
	*/

	db := mahasiswaController.db

	response := responses.Response[string]{}

	insertNewMahasiswaRequest := requests.InsertMahasiswaRequest{}
	if err := c.Bind(&insertNewMahasiswaRequest); err != nil {
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	// selain 0 dan 1 error bad request
	if !(insertNewMahasiswaRequest.Gender == 0 || insertNewMahasiswaRequest.Gender == 1) {
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	// mengambil semua id dan nama hobi yang terdaftar
	hobiRegistered := map[string]int{}
	rows, err := db.Query("SELECT id, nama_hobi from hobi")
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer rows.Close()

	for rows.Next() {
		var hobi string
		var idHobi int
		if err := rows.Scan(&idHobi, &hobi); err != nil {
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
		hobiRegistered[hobi] = idHobi
	}
	if err = rows.Err(); err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	// mengambil semua nama jurusan yang telah terdaftar
	namaJurusanRegistered := []string{}
	rows, err = db.Query("SELECT nama_jurusan from jurusan")
	if err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer rows.Close()

	for rows.Next() {
		jurusan := ""
		if err := rows.Scan(&jurusan); err != nil {
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
		namaJurusanRegistered = append(namaJurusanRegistered, jurusan)
	}
	if err = rows.Err(); err != nil {
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	// TRANSACTION untuk insert ke mahasiswa, jurusan, hobi, mahasiswa_hobi
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

	// insert ke tabel mahasiswa
	result, err := tx.Exec(
		"INSERT INTO mahasiswa (nama, usia, gender, tanggal_registrasi) VALUES (?, ?, ?,?)",
		insertNewMahasiswaRequest.Nama,
		insertNewMahasiswaRequest.Usia,
		insertNewMahasiswaRequest.Gender,
		time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		tx.Rollback()
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	// mengambil id mahasiswa yang baru terdaftar
	IDMahasiswa, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	idSemuaHobi := []int{}
	// insert semua hobi ke tabel hobi
	semuaHobi := insertNewMahasiswaRequest.NamaSemuaHobi
	for _, namaHobi := range semuaHobi {
		_, exist := hobiRegistered[namaHobi]
		if exist {
			idSemuaHobi = append(idSemuaHobi, hobiRegistered[namaHobi])
		} else {
			result, err = tx.Exec("INSERT INTO hobi (nama_hobi) VALUES (?)", strings.ToLower(namaHobi))
			if err != nil {
				tx.Rollback()
				response.Message = types.ERROR_INTERNAL_SERVER
				return c.JSON(http.StatusInternalServerError, response)
			}
			// mengambil id hobi yang baru saja ditmabahkan
			IDHobi, err := result.LastInsertId()
			if err != nil {
				tx.Rollback()
				response.Message = types.ERROR_INTERNAL_SERVER
				return c.JSON(http.StatusInternalServerError, response)
			}
			idSemuaHobi = append(idSemuaHobi, int(IDHobi))
		}
	}

	// insert semua jurusan ke tabel jurusan
	semuaJurusan := insertNewMahasiswaRequest.NamaSemuaJurusan
	for _, namaJurusan := range semuaJurusan {
		if !contains(namaJurusanRegistered, namaJurusan) {
			result, err = tx.Exec("INSERT INTO jurusan (nama_jurusan) VALUES (?)", strings.ToLower(namaJurusan))
			if err != nil {
				tx.Rollback()
				response.Message = types.ERROR_INTERNAL_SERVER
				return c.JSON(http.StatusInternalServerError, response)
			}
		}
	}

	// menambahkan id mahasiswa dan id hobi ke mahasiswa_hobi
	for _, idHobi := range idSemuaHobi {
		_, error := tx.Exec("INSERT INTO mahasiswa_hobi (id_mahasiswa, id_hobi) VALUES (?,?)", IDMahasiswa, idHobi)
		if error != nil {
			tx.Rollback()
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	tx.Commit()

	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}
