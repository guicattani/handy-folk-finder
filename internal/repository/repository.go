package repository

import "github.com/guicattani/handy-folk-finder/internal/models"

type DatabaseRepo interface {
	AllPartners() ([]models.Partner, error)
}
