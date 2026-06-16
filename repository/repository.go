package repository

import (
	"Shannon/shannon"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func (repo *Repository) initTables() {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS PROFILES (
			UserID INTEGER PRIMARY KEY,
			Name TEXT,
			Bio TEXT
		);`
	repo.db.Exec(createTableSQL)
}

func NewRepository(path string) *Repository {
	repo := &Repository{}

	db, err := sql.Open("sqlite", path)

	if err != nil {
		panic("Can't load db")
	}

	repo.db = db

	repo.initTables()

	return repo
}

func (repo *Repository) WriteProfile(profile *shannon.Profile) {
	writeProfileSQL := "INSERT INTO Profiles VALUES (?, ?, ?)"
	_, err := repo.db.Exec(writeProfileSQL, profile.UserID, profile.Name, profile.Bio)
	if err != nil {
		panic("Can't WriteProfile")
	}
}
