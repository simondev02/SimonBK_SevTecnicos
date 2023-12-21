package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Plate            *string `gorm:"type:varchar(20):unique"`
	Vin              *string `gorm:"type:varchar(20):unique"`
	Year             *string `gorm:"type:varchar(20)"`
	Capacity         *int
	SpeedAllowed     *int `gorm:"column:Speedallowed"`
	PurchaseDate     *time.Time
	RegistrationDate *time.Time
	FkCompany        *int
	FkCustomer       *int
	FkBrand          *int
	FkVehType        *int
	FkLine           *int
	FkFuelType       *int
	FkEngineType     *int
	FkAvlDevice      *int `gorm:"type:varchar(20):unique"`
	FkGroupVehicle   *int
	CreditNo         *int
	CaseNo           *int
	VehicleGroup     VehicleGroup `gorm:"foreignKey:FkGroupVehicle"`
	Company          Company      `gorm:"foreignKey:FkCompany"`
	Customer         Customer     `gorm:"foreignKey:FkCustomer"`
	VehicleBrand     VehicleBrand `gorm:"foreignKey:FkBrand"`
	VehicleType      VehicleType  `gorm:"foreignKey:FkVehType"`
	VehicleLine      VehicleLine  `gorm:"foreignKey:FkLine"`
	AvlDevice        AvlDevice    `gorm:"foreignKey:FkAvlDevice"`
	FuelType         FuelType     `gorm:"foreignKey:FkFuelType"`
	EngineType       EngineType   `gorm:"foreignKey:FkEngineType"`
	DeletedBy        *uint
	UpdatedBy        *uint
}
