package repository

import "bri.co.id/nds/approval-service/pkg/domain/entity"

type SupervisorRepository interface {
	GetAllSupervisors() ([]*entity.Supervisor, error)
}
