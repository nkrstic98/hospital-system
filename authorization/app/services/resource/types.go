package resource

import "github.com/google/uuid"

type Assignment struct {
	ActorID     uuid.UUID         `json:"actor_id"`
	Role        string            `json:"role"`
	Permissions map[string]string `json:"permissions"`
}

type JourneyStep struct {
	TransferTime string    `json:"transfer_time"`
	FromTeam     string    `json:"from_team"`
	ToTeam       string    `json:"to_team"`
	FromTeamLead uuid.UUID `json:"from_team_lead"`
	ToTeamLead   uuid.UUID `json:"to_team_lead"`
}

type Resource struct {
	ID          uuid.UUID
	Team        string
	TeamLead    uuid.UUID
	Assignments Assignment
	Journey     []JourneyStep
}
