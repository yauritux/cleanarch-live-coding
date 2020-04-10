package usecase

import (
	"testing"

	repo "bri.co.id/nds/approval-service/pkg/adapter/persistence/inmemory"
	"bri.co.id/nds/approval-service/pkg/domain/entity"

	. "github.com/smartystreets/goconvey/convey"
)

var s []*entity.Supervisor

func setup() {
	s = []*entity.Supervisor{
		&entity.Supervisor{
			ID:     "1",
			Name:   "Yauri",
			Branch: "Jakarta",
		},
	}
}

func TestSpec(t *testing.T) {

	Convey("Feature Approval", t, func() {
		// setup()
		r := repo.NewSupervisorIMem() // collaborator
		uc := NewApproval(r)          // SUT
		Convey("Positive Scenarios", func() {

			Convey("Returns all supervisors", func() {
				res, err := uc.GetSupervisor()
				So(err, ShouldBeEmpty)
				So(res, ShouldNotBeEmpty)
			})
		})
	})
}
