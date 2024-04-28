package models

type User struct {
	UserName string `json:"userName"`
}

type AtcoderUserInfo struct {
	UserName        string    `json:"userName"`
	UserIdentifer   int       `json:"UserIdentifier"`
	Submissions     []int     `json:"submissions"`
	Performances    []int     `json:"performances"`
	Rates           []int     `json:"rates"`
	ContestNames    []string  `json:"contestNames"`
	Similarities    []float64 `json:"similarities"`
	Recommends      []string  `json:"recommends"`
	ProcessedWrongs []string  `json:"processedWrongs"`
}

type Submission struct {
	ID            int64   `json:"id"`
	EpochSecond   int64   `json:"epoch_second"`
	ProblemID     string  `json:"problem_id"`
	ContestID     string  `json:"contest_id"`
	UserID        string  `json:"user_id"`
	Language      string  `json:"language"`
	Point         float64 `json:"point"`
	Length        int     `json:"length"`
	Result        string  `json:"result"`
	ExecutionTime int     `json:"execution_time"`
}

type Performance struct {
	IsRated           bool   `json:"IsRated"`
	Place             int64  `json:"Place"`
	OldRating         int    `json:"OldRating"`
	NewRating         int    `json:"NewRating"`
	Performance       int    `json:"Performance"`
	InnerPerformance  int    `json:"InnerPerformance"`
	ContestScreenName string `json:"ContestScreenName"`
	ContestName       string `json:"ContestName"`
	EndTime           string `json:"EndTime"`
}

type EmbeddingInfo struct {
	Name      string
	Category  string
	Embedding []float64
}

type AlgoData struct {
	ProblemID    string `json:"problem_id"`
	Difficulties string `json:"difficulties"`
	Classifier   string `json:"classifier"`
	URL          string `json:"url"`
}
