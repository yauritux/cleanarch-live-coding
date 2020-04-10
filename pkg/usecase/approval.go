package usecase

import "bri.co.id/nds/approval-service/pkg/domain/repository"

type Approval struct {
	r repository.SupervisorRepository
}

func NewApproval(r repository.SupervisorRepository) *Approval {
	return &Approval{
		r: r,
	}
}

type Supervisor struct {
	ID     string
	Name   string
	Branch string
}

func (a *Approval) GetSupervisor() ([]*Supervisor, error) {
	records, err := a.r.GetAllSupervisors()
	if err != nil {
		return nil, err
	}
	result := make([]*Supervisor, len(records))
	for _, r := range records {
		result = append(result, &Supervisor{
			ID:     r.ID,
			Name:   r.Name,
			Branch: r.Branch,
		})
	}
	return result, nil
}
