package main

import (
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/controllers"
	"GraceClaudia-BackendDeveloper-TechnicalTestJobhunInternship2023/model/types"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "jobhun_local_db",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	e := echo.New()

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println(types.ERROR_CONNECT_TO_DB, err.Error())
		return
	}

	mahasiswaController := controllers.NewMahasiswaController(db)

	// Routes
	// -------------- MAHASISWA --------------
	// Endpoint untuk insert data mahasiswa include jurusan dan hobi
	e.POST("/insertANewMahasiswa", mahasiswaController.InsertANewMahasiswa)

	// Endpoint untuk update data mahasiswa
	e.POST("/updateAMahasiswa", mahasiswaController.UpdateAMahasiswa)

	// Endpoint untuk mendapatkan semua data mahasiswa
	e.GET("/getAllMahasiswa", mahasiswaController.GetAllMahasiswa)

	// 	Endpoint untuk mendapatkan detail data mahasiswa
	e.GET("/getAMahasiswa", mahasiswaController.GetAMahasiswa)

	// Endpoint untuk menghapus data mahasiswa
	e.DELETE("/deleteAMahasiswa", mahasiswaController.DeleteAMahasiswa)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
