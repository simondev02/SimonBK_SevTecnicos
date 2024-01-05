package views

import "time"

type TechnicalService struct {
	Id          uint       `json:"id"`
	Plate       *string    `json:"plate"`
	Vin         *string    `json:"vin"`
	ServiceType *string    `json:"serviceType"`
	Technician  *string    `json:"technician"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	Company     *string    `json:"company"`
	Customer    *string    `json:"customer"`
}
