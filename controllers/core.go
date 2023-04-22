package controllers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type mahasiswaController struct {
	db *sql.DB
}

type MahasiswaController interface {
	InsertANewMahasiswa(c echo.Context) error
	UpdateAMahasiswa(c echo.Context) error
	GetAllMahasiswa(c echo.Context) error
	GetAMahasiswa(c echo.Context) error
	DeleteAMahasiswa(c echo.Context) error
}

func NewMahasiswaController(sqlDatabase *sql.DB) MahasiswaController {
	return &mahasiswaController{db: sqlDatabase}
}
