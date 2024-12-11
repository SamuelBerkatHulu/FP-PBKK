package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// String koneksi tanpa password
	dsn := "root:@tcp(127.0.0.1:3306)/your_jobs?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	// Memastikan koneksi ke database berhasil
	if err := db.Ping(); err != nil {
		log.Fatalf("Tidak dapat terkoneksi ke database: %v", err)
	}

	log.Println("Database terkoneksi")
	DB = db
}
