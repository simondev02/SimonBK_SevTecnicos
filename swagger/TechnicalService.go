package swagger

import "time"

type TechnicalService struct {
	FkVehicle              *uint
	FkTechnicalServiceType *uint
	StartDate              *time.Time
	EndDate                *time.Time
}
