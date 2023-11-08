package schemas

import "time"

type Company struct {
	Name      string
	HoldingID uint
}

type CompanyResponse struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletetedAt time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	HoldingID   uint      `json:"holdingId"`
}
