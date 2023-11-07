package schemas

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	HoldingID uint
	Name      string
}
