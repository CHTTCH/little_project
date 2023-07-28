package patient

import(
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

type Repository[T patient.Patient] interface {
	Save(patient.Patient) error
}
