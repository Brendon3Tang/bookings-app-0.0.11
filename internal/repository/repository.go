// define the interface in repository.go
package repository

import "github.com/tsawler/bookings-app/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
