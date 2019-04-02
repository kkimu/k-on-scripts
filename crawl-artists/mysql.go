package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:pass@tcp(127.0.0.1:3307)/k_on?parseTime=true")

	// テーブルを空に
	//db.Exec("TRUNCATE TABLE artists")
}

func insert(artist Artist) error {
	stmtIns, err := db.Prepare("INSERT INTO artists (id, name, kanaPrefix) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(artist.id, artist.name, artist.kanaPrefix)
	if err != nil {
		return err
	}
	return nil
}

func getArtistByName(name string) (Artist, error) {
	stmtOut, err := db.Prepare("SELECT id FROM artists WHERE name = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close() // Close the statement when we leave main() / the program terminates

	artist := Artist{}
	if err := stmtOut.QueryRow(name).Scan(&artist.id); err != nil {
		return artist, err
	}

	return artist, nil
}
