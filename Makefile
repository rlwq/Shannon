DB := sqlite3
DB_FILE := ./db/file.db

profiles:
	$(DB) $(DB_FILE) "SELECT * FROM Profiles" 
