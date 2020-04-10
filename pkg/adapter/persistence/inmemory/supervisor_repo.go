package inmemory

import (
	"bri.co.id/nds/approval-service/pkg/domain/entity"
)

type SupervisorIMem struct{}

func NewSupervisorIMem() *SupervisorIMem {
	return &SupervisorIMem{}
}

func (*SupervisorIMem) GetAllSupervisors() ([]*entity.Supervisor, error) {
	s := []*entity.Supervisor{
		&entity.Supervisor{
			ID:     "1",
			Name:   "Yauri",
			Branch: "Jakarta",
		},
	}
	return s, nil
}
