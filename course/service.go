package course



import(

	"github.com/robi_a21/Cassiopeia/entity"
)


type CourseService interface{

	AddCourse(course entity.Course) error
	GetCourse() ([]entity.Course,error)

}