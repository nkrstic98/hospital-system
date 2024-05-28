package team

import "hospital-system/authorization/app/services/actor"

type Team struct {
	Name        string
	DisplayName string
	Actors      []actor.Actor
}
