package platons

import (
	"database/sql"
	"log"
)

type PlatonsRepository interface {
	Create(platon Platons) error
	GetPlatonByPostalCode(postalCode int) (*Platons, error)
}

type platonRepository struct {
	db *sql.DB
}

func NewPlatonRepository(db *sql.DB) PlatonsRepository {
	return &platonRepository{db}
}

func (r *platonRepository) Create(platon Platons) error {
	_, err := r.db.Exec("INSERT INTO platons (OriginLatitude, OriginLongitude, OriginPostalCode, Items, Couriers, DestinationLatitude, DestinationLongitude, DestinationPostalCode) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
	platon.OriginLatitude, platon.OriginLongitude, platon.OriginPostalCode, platon.Items, platon.Couriers, platon.DestinationLatitude, platon.DestinationLongitude, platon.DestinationPostalCode)
	if err != nil {
		log.Println("Eror insert data :", err)
	}
	return nil
} 

func (r *platonRepository) GetPlatonByPostalCode(postalCode int) (*Platons, error) {
	row := r.db.QueryRow("SELECT OriginLatitude, OriginLongitude, OriginPostalCode, Items, Couriers, DestinationLatitude, DestinationLongitude, DestinationPostalCode FROM platons WHERE DestinationPostalCode = ? ", postalCode)

	platon := &Platons{}
		err := row.Scan(&platon.OriginLatitude, &platon.OriginLongitude, &platon.OriginPostalCode,  &platon.Items, &platon.Couriers, &platon.DestinationLatitude, &platon.DestinationLongitude, &platon.DestinationPostalCode)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil , nil
		}
		return nil, err
	}
	return platon, nil
}
