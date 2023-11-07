package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Holding struct {
	gorm.Model
	Name      string
	Companies []Company
}

type HoldingResponse struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletetedAt time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
}
