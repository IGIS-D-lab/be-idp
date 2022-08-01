package v2

type Message struct {
	MessageNum  int    `json:"msgNum"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Subthread   int    `json:"subthread"`
	MessageTime string `json:"time"`
	Content     string `json:"content"`
}

type MessageOk struct {
	Message string `json:"message"`
}

const (
	LayOut         = "20060102T15:04:05"
	MessageKey     = "message"
	TestMessageNum = 0
)
