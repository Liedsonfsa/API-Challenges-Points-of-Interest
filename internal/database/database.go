package database

import (
	"errors"

	"github.com/Liedsonfsa/API-Challenges-Points-of-Interest/internal/models"
)

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
		// return nil, err
		return nil, errors.New("erro ao estabelecer a conexão com o banco de dados")
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		// return nil, err
		return nil, errors.New("erro ao executar a query")
	}
	defer rows.Close()

	var pois []models.POI

	for rows.Next() {
		var poi models.POI
		if err = rows.Scan(&poi.ID, &poi.Name, &poi.X, &poi.Y); err != nil {
			// return nil, err
			return nil, errors.New("erro ao scaner algum ponto")
		}

		pois = append(pois, poi)
	}

	return pois, nil
}