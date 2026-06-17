package repo

import (
	"Shannon/shannon"
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func (repo *Repo) initTables() {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS PROFILES (
			UserID INTEGER PRIMARY KEY,
			Name TEXT,
			Bio TEXT
		);`
	repo.db.Exec(createTableSQL)
}

func NewRepo(path string) *Repo {
	repo := &Repo{}

	db, err := sql.Open("sqlite", path)

	if err != nil {
		panic("Can't load db")
	}

	repo.db = db

	repo.initTables()

	return repo
}

func (repo *Repo) CreateProfile(profile shannon.Profile) {
	writeProfileSQL := "INSERT INTO Profiles VALUES (?, ?, ?)"
	_, err := repo.db.Exec(writeProfileSQL, profile.UserID, profile.Name, profile.Bio)
	if err != nil {
		panic("Can't WriteProfile")
	}
}

func (repo *Repo) GetProfiles() []shannon.Profile {
	rows, _ := repo.db.Query("SELECT * FROM Profiles")

	profiles := []shannon.Profile{}
	for rows.Next() {
		profile := shannon.Profile{}
		rows.Scan(profile.UserID, profile.Name, profile.Bio)
		profiles = append(profiles, profile)
	}
	rows.Close()

	return profiles
}

func (repo *Repo) DoesProfileExist(user int64) bool {
	row := repo.db.QueryRow("SELECT UserID FROM Profiles WHERE UserID = ?", user)
	return row != nil
}
