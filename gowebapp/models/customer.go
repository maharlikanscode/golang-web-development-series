package models

import "time"

// Customer model is a master list for all your customers table
type Customer struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"first_name"`          // required
	LastName          string    `json:"last_name"`           // required
	CompanyAddress    string    `json:"company_address"`     // required
	TelNo             string    `json:"tel_no"`              // required
	FaxNo             string    `json:"fax_no"`              // optional
	ContactPersonName string    `json:"contact_person_name"` // optional
	ContactPersonNo   string    `json:"contact_person_no"`   // optional
	CreatedBy         int64     `json:"created_by"`
	CreatedDate       time.Time `json:"created_date"`
	ModifiedBy        int64     `json:"modified_by"`
	ModifiedDate      time.Time `json:"modified_date"`
	DeletedBy         int64     `json:"deleted_by"`
	DeletedDate       time.Time `json:"deleted_date"`
	IsActive          bool      `json:"is_active"`
}
