package repository

import "github.com/guicattani/handy-folk-finder/internal/models"

type DatabaseRepo interface {
	AllPartners() ([]models.Partner, error)
	SpecificPartner(id int) (models.Partner, error)
	CustomerLogin(email string, password string) (models.Customer, error)
	ClosestPartner(customer models.Customer, neededExperience string) ([]models.Partner, error)
}
