package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Plate            *string
	Vin              *string
	Year             *string
	Capacity         *uint
	SpeedAllowed     *uint
	PurchaseDate     *time.Time
	RegistrationDate *time.Time
	FkCompany        *uint
	FkCustomer       *uint
	FkBrand          *uint
	FkVehType        *uint
	FkLine           *uint
	FkFuelType       *uint
	FkEngineType     *uint
	FkAvlDevice      *uint
	FkGroupVehicle   *uint
	VehicleGroup     VehicleGroup `gorm:"foreignKey:FkGroupVehicle"`
	Company          Company      `gorm:"foreignKey:FkCompany"`
	Customer         Customer     `gorm:"foreignKey:FkCustomer"`
	VehicleBrand     VehicleBrand `gorm:"foreignKey:FkBrand"`
	VehicleType      VehicleType  `gorm:"foreignKey:FkVehType"`
	VehicleLine      VehicleLine  `gorm:"foreignKey:FkLine"`
	AvlDevice        AvlDevice    `gorm:"foreignKey:FkAvlDevice"`
	FuelType         FuelType     `gorm:"foreignKey:FkFuelType"`
	EngineType       EngineType   `gorm:"foreignKey:FkEngineType"`
	CreatedBy        *uint
	DeletedBy        *uint
	UpdatedBy        *uint
}
