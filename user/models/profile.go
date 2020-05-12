package models

import "time"

//Profile is a struct representing profile model
type Profile struct {
	ID     string `json:"id"`
	BVN    string `json:"bvn"`
	UserID string `json:"-"`
	User   User   `json:"user"`
	ProfileBio
	ProfileWork
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
}

//ProfileBio is a struct embedded in Profile for user's bio details
type ProfileBio struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Avatar      string `json:"avatar"`
	HomeAddress string `json:"home_address"`
}

//ProfileWork is a struct embedded in Profile for user's work details
type ProfileWork struct {
	EmploymentStatus string `json:"employment_status"` //oneof employed, self-employed, unemployed
	EmployerName     string `json:"employer_name"`
	EmployerAddress  string `json:"employer_address"`
	Designation      string `json:"designation"`
	DateOfEmployment string `json:"date_of_employment"`
}
