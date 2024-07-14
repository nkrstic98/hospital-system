package department

import "hospital-system/server/app/dto"

type Service interface {
	GetDepartments() (map[string]dto.Department, error)
}
