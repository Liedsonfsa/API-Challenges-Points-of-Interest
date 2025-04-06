package database

import "github.com/Liedsonfsa/API-Challenges-Points-of-Interest/internal/models"

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

func GetAllPoints() ([]models.POI, error) {
	query := "SELECT * FROM pois"

	db, err := initDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pois []models.POI

	for rows.Next() {
		var poi models.POI
		if err = rows.Scan(&poi.ID, &poi.Name, &poi.X, &poi.Y); err != nil {
			return nil, err
		}

		pois = append(pois, poi)
	}

	return pois, nil
}

func GetPoints(x, y, d_max uint64) ([]string, error) {
	query := "SELECT * FROM pois"

	db, err := initDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pois []string

	for rows.Next() {
		var poi models.POI
		if err = rows.Scan(&poi.ID, &poi.Name, &poi.X, &poi.Y); err != nil {
			return nil, err
		}

		if poi.WithinTheDistance(x, y, d_max) {
			pois = append(pois, poi.Name)
		}
	}

	return pois, nil
}