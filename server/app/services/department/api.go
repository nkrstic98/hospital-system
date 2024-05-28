package department

type Service interface {
	GetDepartments() (map[string]Department, error)
}
