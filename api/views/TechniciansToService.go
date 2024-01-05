package views

import "time"

type TechniciansToService struct {
	Technician    *string
	Service       uint
	ServiceType   *string
	ServiceStatus *string
	Plate         *string
	Vin           *string
	CreatedAt     *time.Time
}
