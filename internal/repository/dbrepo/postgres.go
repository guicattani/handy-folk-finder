package dbrepo

import (
	"context"
	"fmt"
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

func (m *postgresDBRepo) CustomerLogin(email string, password string) (models.Customer, error) {
	var customer models.Customer

	stmt := fmt.Sprintf(`SELECT * FROM customers WHERE email = '%s' AND password = '%s' LIMIT 1`, email, password)
	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		log.Println(err)
		return customer, err
	}

	err = pgxscan.ScanOne(&customer, rows)

	if err != nil {
		log.Println(err)
	}

	return customer, err
}

func (m *postgresDBRepo) ClosestPartner(customer models.Customer, needed_experience string) ([]models.Partner, error) {
	var partners []models.Partner

	//Spherical Law of Cosines https://www.movable-type.co.uk/scripts/latlong.html#cosine-law
	earth_radius := 6371.0
	stmt := fmt.Sprintf(`
		SELECT id,
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
			FROM (
				SELECT id,
				first_name,
				last_name,
				email,
				COALESCE(rating, -1) as rating,
				address_lat,
				address_lon,
				operating_radius,
				experience,
				created_at,
				updated_at,
				ACOS(
					SIN(RADIANS(%f)) * SIN(RADIANS(partners.address_lat))
					+ COS(RADIANS(%f)) * COS(RADIANS(partners.address_lat))
					* COS(RADIANS(%f) - RADIANS(partners.address_lon))
				) * %f as distance
				FROM partners
				ORDER BY rating DESC, distance ASC
			) AS partners_subquery
			WHERE distance <= partners_subquery.operating_radius
			AND partners_subquery.experience  @> '{%s}'
			`, customer.AddressLat, customer.AddressLat, customer.AddressLon, earth_radius, needed_experience)

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		log.Println(err)
		return partners, err
	}
	defer rows.Close()

	for rows.Next() {
		var partner models.Partner = models.Partner{}

		err = rows.Scan(
			&partner.ID,
			&partner.FirstName,
			&partner.LastName,
			&partner.Email,
			&partner.Rating,
			&partner.AddressLat,
			&partner.AddressLon,
			&partner.OperatingRadius,
			&partner.Experience,
			&partner.CreatedAt,
			&partner.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
			return partners, err
		}
		partners = append(partners, partner)
	}

	return partners, err
}
