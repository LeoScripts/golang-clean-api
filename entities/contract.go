package entities

type StudentUsecaseContract interface {
	List() ([]Student, error)
}
