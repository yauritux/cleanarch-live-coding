package usecase

type ApprovalInputPort interface {
	GetSupervisor() ([]*Supervisor, error)
}
