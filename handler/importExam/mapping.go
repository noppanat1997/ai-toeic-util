package importExam

import "fmt"

func ExecuteMapping(input []Exam) (output GraphQLInput) {
	var beforeExamType float64 = input[0].Type
	var bundleList []ImportExamBundleInput
	var partCount int = 0

	var partTypes = []E_PartType{"PART1_PHOTOGRAPHS",
		"PART2_QUESTION_RESPONSE",
		"PART3_CONVERSATIONS",
		"PART4_TALKS",
		"PART5_INCOMPLETE_SENTENCES",
		"PART6_TEXT_COMPLETION",
		"PART7_READING_COMPREHENSION"}
	var testType E_TestType = Listening
	var questionNo int = 0

	for i, exam := range input {
		fmt.Println(exam.Type, beforeExamType, partCount)

		if partCount > 3 {
			testType = Reading
		}
		switch exam.Haschild {
		case 0:
			var answerList []ImportExamAnswerObjectInput
			for j, correctAns := range exam.Answer.Texts {
				answerList = append(answerList, ImportExamAnswerObjectInput{
					AnswerIndex:    j,
					Text:           correctAns,
					ExplanationAns: "",
					IsCorrect:      true,
				})
			}
			for j, incorrectAns := range exam.Answer.Choices {
				answerList = append(answerList, ImportExamAnswerObjectInput{
					AnswerIndex:    j + len(exam.Answer.Texts),
					Text:           incorrectAns,
					ExplanationAns: "",
					IsCorrect:      false,
				})
			}
			images := []string{}
			if exam.Question.Image != "" {
				images = append(images, exam.Question.Image)
			}
			bundleList = append(bundleList, ImportExamBundleInput{
				BundleNo:    i,
				Description: "",
				TestType:    testType,
				PartType:    partTypes[partCount],
				BundleQuestion: ImportExamBundleBundleQuestionInput{
					Text:   "",
					Images: []string{},
					Sound:  "",
					Hint:   "",
				},
				QuestionList: []ImportExamBundleQuestionInput{
					{
						QuestionNo:  questionNo,
						Explanation: "",
						Tags:        []string{},
						Question: ImportExamBundleBundleQuestionInput{
							Text:   exam.Question.Text,
							Images: images,
							Sound:  exam.Question.Sound,
							Hint:   exam.Question.Hint,
						},
						AnswerList: answerList,
					},
				},
			})
			questionNo += 1
		case 1:
			var questionList []ImportExamBundleQuestionInput
			for _, q := range exam.Childcards {
				images := []string{}
				if q.Question.Image != "" {
					images = append(images, q.Question.Image)
				}
				var answerList []ImportExamAnswerObjectInput
				for j, correctAns := range q.Answer.Texts {
					answerList = append(answerList, ImportExamAnswerObjectInput{
						AnswerIndex:    j,
						Text:           correctAns,
						ExplanationAns: "",
						IsCorrect:      true,
					})
				}
				for j, incorrectAns := range q.Answer.Choices {
					answerList = append(answerList, ImportExamAnswerObjectInput{
						AnswerIndex:    j + len(q.Answer.Texts),
						Text:           incorrectAns,
						ExplanationAns: "",
						IsCorrect:      false,
					})
				}
				questionList = append(questionList, ImportExamBundleQuestionInput{
					QuestionNo:  questionNo,
					Explanation: "",
					Tags:        []string{},
					Question: ImportExamBundleBundleQuestionInput{
						Text:   q.Question.Text,
						Images: images,
						Sound:  q.Question.Sound,
						Hint:   q.Question.Hint,
					},
					AnswerList: answerList,
				})
				questionNo += 1
			}
			images := []string{}
			if exam.Question.Image != "" {
				images = append(images, exam.Question.Image)
			}
			bundleList = append(bundleList, ImportExamBundleInput{
				BundleNo:    i,
				Description: "",
				TestType:    testType,
				PartType:    partTypes[partCount],
				BundleQuestion: ImportExamBundleBundleQuestionInput{
					Text:   exam.Question.Text,
					Images: images,
					Sound:  exam.Question.Sound,
					Hint:   exam.Question.Hint,
				},
				QuestionList: questionList,
			})
		}

		if exam.Type != beforeExamType && partCount < len(partTypes)-1 {
			partCount += 1
		}
		beforeExamType = exam.Type
	}

	output = GraphQLInput{
		Input: ImportExamInput{
			ExamType:    Test,
			Name:        "exam.name",
			Description: "exam.dedescription",
			BundleList:  bundleList,
		},
	}
	return
}
