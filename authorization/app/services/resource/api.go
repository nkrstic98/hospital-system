package resource

import "github.com/google/uuid"

type Service interface {
	AddResource(request Resource) error
	GetResources(ids *[]string) ([]Resource, error)
	ArchiveResource(id uuid.UUID) error
}
