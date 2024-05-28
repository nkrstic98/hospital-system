package department

import "github.com/google/uuid"

type Department struct {
	DisplayName string     `json:"displayName"`
	Physicians  []Employee `json:"physicians"`
	Residents   []Employee `json:"residents"`
	Nurses      []Employee `json:"nurses"`
}

type Employee struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
}
