package dbrepo

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/guicattani/handy-folk-finder/internal/models"
)

func (m *postgresDBRepo) AllPartners() ([]models.Partner, error) {
	var partners []models.Partner

	err := pgxscan.Select(context.Background(), m.DB, &partners, `SELECT id,
																																first_name,
																																last_name,
																																email,
																																COALESCE(rating, -1) as rating,
																																address_lat,
																																address_lon,
																																operating_radius,
																																experience,
																																created_at,
																																updated_at
																																FROM partners
																																ORDER BY rating`)

	if err != nil {
		log.Println(err)
	}

	return partners, err
}

func (m *postgresDBRepo) SpecificPartner(id int) (models.Partner, error) {
	var partner models.Partner

	stmt := fmt.Sprintf(`SELECT id,
											 first_name,
											 last_name,
											 email,
											 COALESCE(rating, -1) as rating,
											 address_lat,
											 address_lon,
											 operating_radius,
											 experience,
											 created_at,
											 updated_at
											 FROM partners
											 WHERE id = %d`, id)
	rows, _ := m.DB.Query(context.Background(), stmt)
	err := pgxscan.ScanOne(&partner, rows)

	if err != nil {
		log.Println(err)
	}

	return partner, err
}
}
