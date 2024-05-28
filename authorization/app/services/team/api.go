package team

type Service interface {
	GetTeam(team string) (Team, error)
	GetTeams() ([]Team, error)
}
