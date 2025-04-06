package database


func RegisterPOI(name string, x, y uint64) error {
	query := "INSERT INTO pois (name, x, y) VALUES (?, ?, ?)"

	db, err := initDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, x, y)

	return err
}