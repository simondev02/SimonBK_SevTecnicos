package views

type TechnicalServiceStatus struct {
	ID   uint
	Name *string `gorm:"unique"`
}
