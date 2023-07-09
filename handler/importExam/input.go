package importExam

type Exam struct {
	Id              string   `json:"_id"`
	Parentid        string   `json:"parentId"`
	Haschild        int      `json:"hasChild"`
	Question        Question `json:"question"`
	Answer          Answer   `json:"answer"`
	Code            string   `json:"code"`
	Status          int      `json:"status"`
	Lastupdate      float64  `json:"lastUpdate"`
	Difficultylevel int      `json:"difficultyLevel"`
	Orderindex      float64  `json:"orderIndex"`
	Type            float64  `json:"type"`
	Setting         Setting  `json:"setting"`
	Childcards      []Exam   `json:"childCards"`
	Isrender        bool     `json:"isRender"`
	Tags            []string `json:"tags"`
	Games           []string `json:"games"`
	Pairs           []string `json:"pairs"`
	Trappairs       []string `json:"trapPairs"`
}

type Question struct {
	Text  string `json:"text"`
	Image string `json:"image"`
	Sound string `json:"sound"`
	Hint  string `json:"hint"`
}

type Answer struct {
	Texts   []string `json:"texts"`
	Choices []string `json:"choices"`
	Image   string   `json:"image"`
	Sound   string   `json:"sound"`
	Hint    string   `json:"hint"`
}

type Setting struct {
	ShuffleAnswer     int    `json:"shuffleAnswer"`
	Additional        string `json:"additional"`
	MaxPoint          int    `json:"maxPoint"`
	PointType         int    `json:"pointType"`
	SplitPointAnswers bool   `json:"splitPointAnswers"`
}
