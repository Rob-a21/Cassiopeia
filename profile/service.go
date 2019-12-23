package profile

import (
<<<<<<< HEAD
	"github.com/Rob_a21/Cassiopeia/entity"
=======
	"github.com/robi_a21/Cassiopeia/entity"
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
)

type ProfileService interface {
	Students() ([]entity.Student, error)

	Student(id int) (entity.Student, error)
}
