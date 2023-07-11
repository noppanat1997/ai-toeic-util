package importExam

type GraphQLInput struct {
	Input ImportExamInput `json:"input"`
}

type ImportExamInput struct {
	ExamType    E_ExamType              `json:"examType"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	BundleList  []ImportExamBundleInput `json:"bundleList"`
}

type ImportExamBundleInput struct {
	BundleNo       int                                 `json:"bundleNo"`
	Description    string                              `json:"description"`
	TestType       E_TestType                          `json:"testType"`
	PartType       E_PartType                          `json:"partType"`
	BundleQuestion ImportExamBundleBundleQuestionInput `json:"bundleQuestion"`
	QuestionList   []ImportExamBundleQuestionInput     `json:"questionList"`
}

type ImportExamBundleBundleQuestionInput struct {
	Text   string   `json:"text"`
	Images []string `json:"images"`
	Sound  string   `json:"sound"`
	Hint   string   `json:"hint"`
}

type ImportExamBundleQuestionInput struct {
	QuestionNo  int                                 `json:"questionNo"`
	Explanation string                              `json:"explanation"`
	Tags        []string                            `json:"tags"`
	Question    ImportExamBundleBundleQuestionInput `json:"question"`
	AnswerList  []ImportExamAnswerObjectInput       `json:"answerList"`
}

type ImportExamAnswerObjectInput struct {
	AnswerIndex    int    `json:"answerIndex"`
	Text           string `json:"text"`
	ExplanationAns string `json:"explanationAns"`
	IsCorrect      bool   `json:"isCorrect"`
}
