package database

import (
	"github.com/jeffrpowell/hellogo/internal/constants"
	"github.com/lib/pq"
)

func GetGradient(name string) (constants.ColorGradient, error) {
	db := getDatabaseConnection()
	defer db.Close()
	row := db.QueryRow("SELECT Id, Name, Colors FROM "+constants.DB_TABLE_COLORGRADIENT+" WHERE Name = $1", name)
	var gradient constants.ColorGradient
	err := row.Scan(&gradient.Id, &gradient.Name, pq.Array(&gradient.Colors))
	if err != nil {
		return constants.ColorGradient{}, err
	}
	return gradient, nil
}
