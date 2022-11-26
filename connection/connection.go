package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4" // Install pgx postgresql (go get -u github.com/jackc/pgx/v5)
)

// buat pointer ke pgx agar dapat diakses secara global
var Conn *pgx.Conn

// Koneksi database
func DatabaseConnect() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	var err error
	databaseUrl := "postgres://postgres:root@localhost:5432/blogs"
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Koneksi ke database gagal: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Koneksi ke database berhasil")
}

// Pada code diatas, kita memanggil package pgx dan melakukan konfigurasi untuk koneksi database.
