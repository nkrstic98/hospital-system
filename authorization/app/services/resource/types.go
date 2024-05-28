package resource

import "github.com/google/uuid"

type Resource struct {
	ID       uuid.UUID
	Team     string
	TeamLead uuid.UUID
}
