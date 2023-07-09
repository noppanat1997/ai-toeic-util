package importExam

type E_ExamType string
type E_TestType string
type E_PartType string

const (
	Test    E_ExamType = "TEST"
	Lecture            = "LECTURE"
)

const (
	Listening E_TestType = "LISTENING"
	Reading              = "READING"
)
